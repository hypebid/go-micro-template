package main

import (
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"

	"github.com/hypebid/go-micro-template/internal/config"
	"github.com/hypebid/go-micro-template/internal/rpc"
	"github.com/hypebid/go-micro-template/internal/rpc/pb"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.NewServiceConfig()
	if err != nil {
		log.Printf("Error initializing service config: %v", err)
		return
	}

	logOpts := []grpc_logrus.Option{}
	recovOpts := []grpc_recovery.Option{}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", c.Constants.Port))
	if err != nil {
		c.Log.Fatalf("Failed to listen: %v\n", err)
		return
	}

	grpcServer := grpc.NewServer(
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(logrus.NewEntry(c.Log), logOpts...),
			grpc_recovery.StreamServerInterceptor(recovOpts...)),
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(c.Log), logOpts...),
			grpc_recovery.UnaryServerInterceptor(recovOpts...)),
	)

	pb.RegisterServiceNameServer(grpcServer, &rpc.Server{Config: c})

	c.Log.Printf("Server listening on %v", c.Constants.Port)

	if err := grpcServer.Serve(lis); err != nil {
		c.Log.Fatalf("failed to serve: %v\n", err)
	}
}
