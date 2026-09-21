package services

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/webtor-io/claims-provider/models"
	pb "github.com/webtor-io/claims-provider/proto"
	"github.com/webtor-io/lazymap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

func TestLookupOutcome(t *testing.T) {
	cases := []struct {
		name   string
		claims *models.Claims
		err    error
		want   string
	}{
		{"row", &models.Claims{TierID: 1}, nil, outcomeOK},
		{"no rows", nil, pg.ErrNoRows, outcomeNotFound},
		{"wrapped no rows", nil, wrap(pg.ErrNoRows), outcomeNotFound},
		{"nil claims nil err", nil, nil, outcomeNotFound},
		{"db error", nil, errors.New("connection refused"), outcomeError},
		{"timeout", nil, context.DeadlineExceeded, outcomeError},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := lookupOutcome(c.claims, c.err); got != c.want {
				t.Fatalf("lookupOutcome(%v, %v) = %q, want %q", c.claims, c.err, got, c.want)
			}
		})
	}
}

type wrapped struct{ err error }

func (w wrapped) Error() string { return "wrapped: " + w.err.Error() }
func (w wrapped) Unwrap() error { return w.err }
func wrap(err error) error      { return wrapped{err} }

// A cache miss is counted once per source with the outcome of the fetch; a
// cache hit is not a lookup and must not be counted.
func TestStoreLookupMetrics(t *testing.T) {
	st := &Store{LazyMap: lazymap.New[*models.Claims](&lazymap.Config{Concurrency: 1, Expire: time.Minute, ErrorExpire: time.Minute, Capacity: 10})}
	st.fetch = func(ctx context.Context, email string) (*models.Claims, error) {
		return &models.Claims{TierID: 1}, nil
	}
	st.fetchByPatreonID = func(ctx context.Context, patreonID string) (*models.Claims, error) {
		return nil, errors.New("boom")
	}

	emailOK := counterValue(t, claimsLookupsTotal.WithLabelValues(lookupSourceEmail, outcomeOK))
	patreonErr := counterValue(t, claimsLookupsTotal.WithLabelValues(lookupSourcePatreonID, outcomeError))

	if _, err := st.GetByEmail(context.Background(), "a@example.com"); err != nil {
		t.Fatal(err)
	}
	// Second call on the same key is served from the cache.
	if _, err := st.GetByEmail(context.Background(), "a@example.com"); err != nil {
		t.Fatal(err)
	}
	if _, err := st.GetByPatreonID(context.Background(), "pat_1"); err == nil {
		t.Fatal("expected error")
	}

	if got := counterValue(t, claimsLookupsTotal.WithLabelValues(lookupSourceEmail, outcomeOK)) - emailOK; got != 1 {
		t.Fatalf("email/ok lookups = %v, want 1 (cache hit must not count)", got)
	}
	if got := counterValue(t, claimsLookupsTotal.WithLabelValues(lookupSourcePatreonID, outcomeError)) - patreonErr; got != 1 {
		t.Fatalf("patreon_id/error lookups = %v, want 1", got)
	}
}

func counterValue(t *testing.T, c prometheus.Counter) float64 {
	t.Helper()
	return testutil.ToFloat64(c)
}

// The gRPC server is instrumented end to end: a call over the wire lands in
// grpc_server_handled_total with the status code the handler returned.
func TestGRPCHandledTotal(t *testing.T) {
	st := &Store{LazyMap: lazymap.New[*models.Claims](&lazymap.Config{Concurrency: 1, Expire: time.Minute, ErrorExpire: time.Minute, Capacity: 10})}
	st.fetch = func(ctx context.Context, email string) (*models.Claims, error) {
		if email == "known@example.com" {
			return &models.Claims{TierID: 1, TierName: "bronze"}, nil
		}
		return nil, nil
	}
	g := &GRPC{store: st}
	gs := g.newServer()
	lis := bufconn.Listen(1 << 20)
	go func() { _ = gs.Serve(lis) }()
	t.Cleanup(gs.Stop)

	conn, err := grpc.NewClient("passthrough:///bufconn",
		grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) { return lis.DialContext(ctx) }),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = conn.Close() })
	cl := pb.NewClaimsProviderClient(conn)

	okBefore := handledTotal(t, codes.OK)
	nfBefore := handledTotal(t, codes.NotFound)

	if _, err := cl.Get(context.Background(), &pb.GetRequest{Email: "known@example.com"}); err != nil {
		t.Fatalf("Get(known): %v", err)
	}
	_, err = cl.Get(context.Background(), &pb.GetRequest{Email: "nobody@example.com"})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("Get(unknown) code = %v, want NotFound", status.Code(err))
	}

	if got := handledTotal(t, codes.OK) - okBefore; got != 1 {
		t.Fatalf("handled_total{OK} delta = %v, want 1", got)
	}
	if got := handledTotal(t, codes.NotFound) - nfBefore; got != 1 {
		t.Fatalf("handled_total{NotFound} delta = %v, want 1", got)
	}
}

// handledTotal reads grpc_server_handled_total for ClaimsProvider/Get with
// the given code off the default registry — the same view /metrics serves.
func handledTotal(t *testing.T, code codes.Code) float64 {
	t.Helper()
	families, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"grpc_service": "ClaimsProvider",
		"grpc_method":  "Get",
		"grpc_type":    "unary",
		"grpc_code":    code.String(),
	}
	for _, f := range families {
		if f.GetName() != "grpc_server_handled_total" {
			continue
		}
	next:
		for _, m := range f.GetMetric() {
			for _, lp := range m.GetLabel() {
				if v, ok := want[lp.GetName()]; ok && v != lp.GetValue() {
					continue next
				}
			}
			return m.GetCounter().GetValue()
		}
	}
	return 0
}
