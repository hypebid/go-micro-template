package tests

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"testing"

	"github.com/hypebid/go-micro-template/internal/config"
	"github.com/hypebid/go-micro-template/internal/rpc"
	"github.com/hypebid/go-micro-template/internal/rpc/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

func createHash(md metadata.MD) (string, error) {
	var hmac_message string
	for _, v := range []string{"rpc-method", "service-name", "hypebid-noauth", "hypebid-nohash"} {
		if len(md.Get(v)) == 0 {
			log.Printf("rpc request does not contain this metadata: %v", v)
			return "", errors.New("rpc request does not contain right metadata")
		}
		hmac_message = fmt.Sprintf("%v%v", hmac_message, md.Get(v)[0])
	}
	log.Println("hmac message created")
	mac := hmac.New(sha256.New, []byte("secretstringvalue"))
	mac.Write([]byte(hmac_message))
	expectedMAC := mac.Sum(nil)

	return string(expectedMAC), nil
}

func TestHealthCheckWithNoHash(t *testing.T) {
	md := metadata.Pairs(
		"rpc-method", "healthCheck",
		"service-name", "testService",
		"hypebid-noauth", "false",
		"hypebid-nohash", "true",
		"hypebid-hash", "baldjfasdkfjkjsd",
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
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

	// asserts
	assert.Equal(t, os.Getenv("SERVICE_NAME"), resp.ServiceName, "service name should match")
}

func TestHealthCheckWithHash_Negative(t *testing.T) {
	md := metadata.Pairs(
		"rpc-method", "healthCheck",
		"service-name", "testService",
		"hypebid-noauth", "false",
		"hypebid-nohash", "false",
		"hypebid-hash", "baldjfasdkfjkjsd",
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewServiceNameClient(conn)
	resp, err := client.HealthCheck(ctx, &pb.HealthRequest{Message: "testing healthcheck"})

	// asserts
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "auth issue")
}

func TestHealthCheckWithHash_Positive(t *testing.T) {
	md := metadata.Pairs(
		"rpc-method", "healthCheck",
		"service-name", "testService",
		"hypebid-noauth", "false",
		"hypebid-nohash", "false",
	)
	hash, err := createHash(md)
	if err != nil {
		t.Fatalf("failed to make hash: %v", err)
	}
	md.Append("hypebid-hash-bin", hash)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewServiceNameClient(conn)
	resp, err := client.HealthCheck(ctx, &pb.HealthRequest{Message: "testing healthcheck"})
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}

	// asserts
	assert.NotNil(t, resp)
}
