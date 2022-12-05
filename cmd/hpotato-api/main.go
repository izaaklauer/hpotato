package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/izaaklauer/hpotato/config"
	hpotatov1 "github.com/izaaklauer/hpotato/gen/proto/go/hpotato/v1"
	"github.com/izaaklauer/hpotato/server"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("starting hpotato.......")
	defer fmt.Println("hpotato exiting!")

	err := serve()
	if err != nil {
		log.Fatalf("failed serving\n%s", err)
	}
}

func serve() error {
	var c config.Config

	configPath := os.Getenv("CONFIG_PATH")
	if configPath != "" {
		var err error
		c, err = config.GetConfig(configPath)
		if err != nil {
			return errors.Wrapf(err, "failed to load config from %q", configPath)
		}
	} else {
		// Erroring on no config would be totally valid
		// return errors.New("Environment variable CONFIG_PATH is unset")

		c = config.DefaultConfig()
	}

	// Start the service
	hpotatoServer, err := server.NewHpotatoServer(c.Hpotato)
	if err != nil {
		return errors.Wrapf(err, "failed to start hpotato server")
	}

	listener, err := net.Listen("tcp", c.Server.BindAddr)
	if err != nil {
		return errors.Wrapf(err, "failed to listen on %s", c.Server.BindAddr)
	}
	grpcServer := grpc.NewServer()

	hpotatov1.RegisterHpotatoServiceServer(grpcServer, hpotatoServer)
	reflection.Register(grpcServer)

	log.Printf("Serving on %q", c.Server.BindAddr)
	if err := grpcServer.Serve(listener); err != nil {
		return errors.Wrapf(err, "gRPC server exited")
	}

	return nil
}
