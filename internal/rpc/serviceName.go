package rpc

import (
	"context"
	"log"

	"github.com/hypebid/go-micro-template/internal/rpc/pb"
)

type Server struct {
	pb.UnsafeServiceNameServer
}

func (s *Server) HealthCheck(ctx context.Context, req *pb.HealthRequest) (*pb.HealthStatus, error) {
	log.Printf("Received: %v", req.GetMessage())

	return &pb.HealthStatus{
		TransactionId:  "00DSKDJF882k2kd",
		ServiceName:    "ServiceName",
		ReleaseDate:    "8/4/2021",
		ReleaseSlug:    "0bsg54dkfjs9292kdkgjg93932",
		ReleaseVersion: "v1.0.1",
		DatabaseOnline: false,
		Message:        req.GetMessage(),
	}, nil
}
