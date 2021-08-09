package rpc

import (
	"context"

	"github.com/hypebid/go-micro-template/internal/config"
	"github.com/hypebid/go-micro-template/internal/rpc/pb"
)

type Server struct {
	pb.UnsafeServiceNameServer
	Config *config.Config
}

func (s *Server) HealthCheck(ctx context.Context, req *pb.HealthRequest) (*pb.HealthStatus, error) {
	s.Config.Log.Printf("Received: %v", req.GetMessage())

	// ping db
	dbOnline := false
	ping := s.Config.Psql.DB.Raw("SELECT * FROM information_schema.information_schema_catalog_name;")
	if ping.Error == nil {
		dbOnline = true
	}

	return &pb.HealthStatus{
		TransactionId:  "00DSKDJF882k2kd",
		ServiceName:    s.Config.Constants.ServiceName,
		ReleaseDate:    s.Config.Constants.ReleaseDate,
		ReleaseSlug:    s.Config.Constants.ReleaseSlug,
		ReleaseVersion: s.Config.Constants.ReleaseVersion,
		DatabaseOnline: dbOnline,
		Message:        req.GetMessage(),
	}, nil
}
