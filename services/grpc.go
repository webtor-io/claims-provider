package services

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	pb "github.com/webtor-io/claims-provider/proto"
	"google.golang.org/grpc"
)

const (
	grpcHostFlag = "grpc-host"
	grpcPortFlag = "grpc-port"
)

func RegisterGRPCFlags(f []cli.Flag) []cli.Flag {
	return append(f,
		cli.StringFlag{
			Name:   grpcHostFlag,
			Usage:  "grpc listening host",
			Value:  "",
			EnvVar: "GRPC_HOST",
		},
		cli.IntFlag{
			Name:   grpcPortFlag,
			Usage:  "grpc listening port",
			Value:  50051,
			EnvVar: "GRPC_PORT",
		},
	)
}

type GRPC struct {
	pb.UnimplementedClaimsProviderServer
	host  string
	port  int
	ln    net.Listener
	store *Store
}

// NewGRPC returns a new GRPC server
func NewGRPC(c *cli.Context, s *Store) *GRPC {
	// Build and return a new GRPC server
	return &GRPC{
		// Host and port to listen on
		host: c.String(grpcHostFlag),
		port: c.Int(grpcPortFlag),
		// Store to get claims from
		store: s,
	}
}

// Serve starts the gRPC server
func (s *GRPC) Serve() error {
	// Build address
	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	// Listen to TCP connection
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrap(err, "failed to listen to tcp connection")
	}

	// Save the listener
	s.ln = listener

	// Create a new gRPC server with a logger interceptor
	serverOpts := []grpc.ServerOption{
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			start := time.Now()
			resp, err := handler(ctx, req)
			log.WithFields(log.Fields{
				"method": info.FullMethod,
				"took":   time.Since(start),
			}).Info("finished unary call")
			if err != nil {
				log.WithError(err).Error("error in unary call")
			}
			return resp, err
		}),
	}
	gs := grpc.NewServer(serverOpts...)

	// Register the service
	pb.RegisterClaimsProviderServer(gs, s)

	// Log the start message
	log.Infof("serving GRPC at %v", addr)

	// Start the server
	return gs.Serve(listener)
}

// Close stops the gRPC server and releases the listener
func (s *GRPC) Close() error {
	// Close the listener if it's open
	if s.ln != nil {
		// Close the listener
		return s.ln.Close()
	}
	// No error if the listener is not open
	return nil
}

// Get gets claims for a user
func (s *GRPC) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	// Get claims from store
	c, err := s.store.Get(req.Email)
	if err != nil {
		return nil, err
	}
	// Build GRPC response
	return &pb.GetResponse{
		Context: &pb.Context{ // Context is a part of claim response
			Tier: &pb.Tier{ // Tier is a part of context
				Id:   c.TierID,   // Tier ID
				Name: c.TierName, // Tier name
			},
		},
		Claims: &pb.Claims{ // Claims is a part of claim response
			Connection: &pb.Connection{ // Connection is a part of claims
				Rate: c.DownloadRate, // Download rate
			},
			Embed: &pb.Embed{ // Embed is a part of claims
				NoAds: c.EmbedNoAds, // No ads
			},
			Site: &pb.Site{ // Site is a part of claims,
				NoAds: c.SiteNoAds, // No ads
			},
		},
	}, nil
}
