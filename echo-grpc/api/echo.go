// The echo-grpc service came from GCP's repo:
// https://github.com/GoogleCloudPlatform/grpc-gke-nlb-tutorial

package api

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Server for the Echo gRPC API
type Server struct{}

// Echo the content of the request
func (s *Server) Echo(ctx context.Context, in *EchoRequest) (*EchoResponse, error) {
	log.Printf("Handling Echo request [%v] with context %v", in, ctx)
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Unable to get hostname %v", err)
		hostname = ""
	}
	grpc.SendHeader(ctx, metadata.Pairs("hostname", hostname))
	content := in.Content + " hostname: " + hostname
	return &EchoResponse{Content: content}, nil
}
