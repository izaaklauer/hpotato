package server

import (
	"context"
	"log"

        "github.com/izaaklauer/hpotato/config"
	hpotatov1 "github.com/izaaklauer/hpotato/gen/proto/go/hpotato/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type HpotatoServer struct {
	hpotatov1.UnimplementedHpotatoServiceServer

	config config.Hpotato
}

// NewHpotatoServer initializes a new server from config
func NewHpotatoServer(config config.Hpotato) (*HpotatoServer, error) {
	// Server-specific initialization, like DB clients, goes here.

	server := HpotatoServer{
		config: config,
	}

	return &server, nil
}

func (s * HpotatoServer) HelloWorld(
	ctx context.Context,
	req *hpotatov1.HelloWorldRequest,
) (*hpotatov1.HelloWorldResponse, error) {
	log.Printf("HelloWorld request with message %q", req.Message)

	resp := &hpotatov1.HelloWorldResponse{
		RequestMessage: req.Message,
		ConfigMessage:  s.config.HelloWorldMessage,
		Now:            timestamppb.Now(),
	}

	return resp, nil
}
