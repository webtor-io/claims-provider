package services

import (
	"errors"
	"time"

	"github.com/go-pg/pg/v10"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/webtor-io/claims-provider/models"
)

// Lookup sources: how the caller identified the user. Values are fixed here,
// not taken from the request, so the label set stays bounded.
const (
	lookupSourceEmail     = "email"
	lookupSourcePatreonID = "patreon_id"
)

// Lookup outcomes. not_found is a legitimate answer (a user with no
// membership row), so it is kept apart from error — otherwise a wave of
// anonymous lookups would look like a database incident, and a real one
// would hide inside the noise.
const (
	outcomeOK       = "ok"
	outcomeNotFound = "not_found"
	outcomeError    = "error"
)

var (
	// grpcMetrics is the standard gRPC server instrumentation
	// (grpc_server_started_total, grpc_server_handled_total{grpc_code},
	// grpc_server_handling_seconds, msg counters). One instance shared by
	// every server this process builds, registered once on the default
	// registry that cs.NewProm exposes on /metrics.
	//
	// Bucket ceiling of 30s: the store's DB timeout is 5s per lookup and a
	// request can do two, so the histogram must resolve well past 10s to
	// show a stuck database instead of clipping into +Inf.
	grpcMetrics = grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10, 30}),
		),
	)

	// claimsLookupsTotal counts backing-store lookups by how the identity
	// was resolved and what came back. Counted at the cache-miss layer:
	// the in-process cache answers most calls, and what this metric is for
	// is the traffic that actually reaches the database — its error rate
	// is the money-path signal, and its share of gRPC calls is the cache
	// hit ratio.
	claimsLookupsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "claims_lookups_total",
		Help: "Claims lookups against the backing store (cache misses), by identity source and outcome.",
	}, []string{"source", "outcome"})

	// claimsLookupSeconds is the latency of the same lookups. Buckets stop
	// at 10s: the per-lookup DB timeout is 5s, so anything beyond is the
	// timeout firing, and one bucket past it is enough to see that.
	claimsLookupSeconds = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "claims_lookup_seconds",
		Help:    "Latency of claims lookups against the backing store, by identity source.",
		Buckets: []float64{0.001, 0.0025, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
	}, []string{"source"})
)

func init() {
	// promauto registered the two vectors above; the gRPC collector is a
	// plain Collector and has to be registered by hand.
	prometheus.MustRegister(grpcMetrics)
}

// lookupOutcome maps a store lookup result onto a metric label.
//
// A nil claims with a nil error is not_found too: the injected fetchers used
// in tests answer that way, and the Get handler treats it exactly like a
// missing row, so the metric must agree with the handler.
func lookupOutcome(claims *models.Claims, err error) string {
	switch {
	case errors.Is(err, pg.ErrNoRows):
		return outcomeNotFound
	case err != nil:
		return outcomeError
	case claims == nil:
		return outcomeNotFound
	}
	return outcomeOK
}

// instrumentLookup wraps a cache-miss builder so the lookup is counted and
// timed. It wraps the builder rather than the DB call so every miss is
// measured the same way whichever fetcher is installed.
func instrumentLookup(source string, fetch func() (*models.Claims, error)) func() (*models.Claims, error) {
	return func() (*models.Claims, error) {
		start := time.Now()
		claims, err := fetch()
		claimsLookupSeconds.WithLabelValues(source).Observe(time.Since(start).Seconds())
		claimsLookupsTotal.WithLabelValues(source, lookupOutcome(claims, err)).Inc()
		return claims, err
	}
}
