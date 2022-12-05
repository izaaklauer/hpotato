package server

import (
	"context"
	"log"

        "github.com/izaaklauer/%%wp_project%%/config"
	%%wp_project%%v1 "github.com/izaaklauer/%%wp_project%%/gen/proto/go/%%wp_project%%/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type %%Wp_project%%Server struct {
	%%wp_project%%v1.Unimplemented%%Wp_project%%ServiceServer

	config config.%%Wp_project%%
}

// New%%Wp_project%%Server initializes a new server from config
func New%%Wp_project%%Server(config config.%%Wp_project%%) (*%%Wp_project%%Server, error) {
	// Server-specific initialization, like DB clients, goes here.

	server := %%Wp_project%%Server{
		config: config,
	}

	return &server, nil
}

func (s * %%Wp_project%%Server) HelloWorld(
	ctx context.Context,
	req *%%wp_project%%v1.HelloWorldRequest,
) (*%%wp_project%%v1.HelloWorldResponse, error) {
	log.Printf("HelloWorld request with message %q", req.Message)

	resp := &%%wp_project%%v1.HelloWorldResponse{
		RequestMessage: req.Message,
		ConfigMessage:  s.config.HelloWorldMessage,
		Now:            timestamppb.Now(),
	}

	return resp, nil
}
