package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_reqAuth "github.com/hypebid/go-kit/grpc/middleware/auth"
	grpc_reqId "github.com/hypebid/go-kit/grpc/middleware/transactionId"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	"github.com/hypebid/go-micro-template/internal/config"
	"github.com/hypebid/go-micro-template/internal/rpc"
	"github.com/hypebid/go-micro-template/internal/rpc/pb"
	"google.golang.org/grpc"
)

func metrics() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2111", nil)
}

func main() {
	c, err := config.NewServiceConfig()
	if err != nil {
		log.Printf("Error initializing service config: %v", err)
		return
	}

	logOpts := []grpc_logrus.Option{}
	recovOpts := []grpc_recovery.Option{}
	reqAuthOpts := grpc_reqAuth.Options{
		HashSecret: c.Constants.HashSecret,
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", c.Constants.Port))
	if err != nil {
		c.Log.Fatalf("Failed to listen: %v\n", err)
		return
	}

	grpcServer := grpc.NewServer(
		grpc_middleware.WithStreamServerChain(
			grpc_reqId.StreamServerInterceptor(c.Log),
			grpc_reqAuth.StreamServerInterceptor(c.Log),
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(logrus.NewEntry(c.Log), logOpts...),
			grpc_recovery.StreamServerInterceptor(recovOpts...)),
		grpc_middleware.WithUnaryServerChain(
			grpc_reqId.UnaryServerInterceptor(c.Log),
			grpc_reqAuth.UnaryServerInterceptor(c.Log, reqAuthOpts),
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(c.Log), logOpts...),
			grpc_recovery.UnaryServerInterceptor(recovOpts...)),
	)

	go metrics()

	pb.RegisterServiceNameServer(grpcServer, &rpc.Server{Config: c})

	c.Log.Printf("Server listening on %v", c.Constants.Port)

	if err := grpcServer.Serve(lis); err != nil {
		c.Log.Fatalf("failed to serve: %v\n", err)
	}
}
