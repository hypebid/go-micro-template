package main

import (
	"context"
	"fmt"
	"net"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	log "github.com/sirupsen/logrus"

	"github.com/hypebid/go-micro-template/internal/grpc/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements grpc helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v\n", in.GetName())

	return &pb.HelloReply{Message: fmt.Sprintf("Hello %v\n", in.GetName())}, nil
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.InfoLevel)
}

func main() {
	// var logrusLogger *log.Logger

	// logrusEntry := log.NewEntry(logrusLogger)
	opts := []grpc_logrus.Option{}

	// grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	grpcServer := grpc.NewServer(
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(log.NewEntry(log.New()), opts...)),
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(log.NewEntry(log.New()), opts...)),
	)

	pb.RegisterGreeterServer(grpcServer, &server{})

	log.Println("Server listening on 8080")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
