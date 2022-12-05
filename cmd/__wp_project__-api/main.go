package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/izaaklauer/%%wp_project%%/config"
	%%wp_project%%v1 "github.com/izaaklauer/%%wp_project%%/gen/proto/go/%%wp_project%%/v1"
	"github.com/izaaklauer/%%wp_project%%/server"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("starting %%wp_project%%.......")
	defer fmt.Println("%%wp_project%% exiting!")

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
	%%wp_project%%Server, err := server.New%%Wp_project%%Server(c.%%Wp_project%%)
	if err != nil {
		return errors.Wrapf(err, "failed to start %%wp_project%% server")
	}

	listener, err := net.Listen("tcp", c.Server.BindAddr)
	if err != nil {
		return errors.Wrapf(err, "failed to listen on %s", c.Server.BindAddr)
	}
	grpcServer := grpc.NewServer()

	%%wp_project%%v1.Register%%Wp_project%%ServiceServer(grpcServer, %%wp_project%%Server)
	reflection.Register(grpcServer)

	log.Printf("Serving on %q", c.Server.BindAddr)
	if err := grpcServer.Serve(listener); err != nil {
		return errors.Wrapf(err, "gRPC server exited")
	}

	return nil
}
