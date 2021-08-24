package tests

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/hypebid/go-micro-template/internal/config"
	"github.com/hypebid/go-micro-template/internal/rpc"
	"github.com/hypebid/go-micro-template/internal/rpc/pb"
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

	lis = bufconn.Listen(bufSize)
	_, grpcServer, err := rpc.RpcSetup(c)
	if err != nil {
		log.Fatalf("failed to setup rpc server")
	}

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
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
