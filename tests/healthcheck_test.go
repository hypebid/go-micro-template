package tests

import (
	"context"
	"log"
	"net"
	"strings"
	"testing"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_reqAuth "github.com/hypebid/go-kit/grpc/middleware/auth"
	grpc_reqId "github.com/hypebid/go-kit/grpc/middleware/transactionId"
	"github.com/hypebid/go-micro-template/internal/config"
	"github.com/hypebid/go-micro-template/internal/rpc"
	"github.com/hypebid/go-micro-template/internal/rpc/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	c, err := config.NewServiceConfig()
	if err != nil {
		log.Fatalf("failed to create config")
	}
	logOpts := []grpc_logrus.Option{}
	recovOpts := []grpc_recovery.Option{}
	reqAuthOpts := grpc_reqAuth.Options{
		HashSecret:      c.Constants.HashSecret,
		MetadataKeyList: strings.Split(c.Constants.MetadataKeyList, ","),
		MetadataHashKey: c.Constants.MetadataHashKey,
	}
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer(
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
	pb.RegisterServiceNameServer(s, &rpc.Server{Config: c})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestHealthCheck(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewServiceNameClient(conn)
	resp, err := client.HealthCheck(ctx, &pb.HealthRequest{Message: "testing healthcheck"})
	if err != nil {
		t.Fatalf("health check failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
}
