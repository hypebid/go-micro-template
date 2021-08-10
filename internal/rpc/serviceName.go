package rpc

import (
	"context"

	grpc_reqId "github.com/hypebid/go-kit/grpc/middleware/transactionId"
	"github.com/hypebid/go-micro-template/internal/config"
	"github.com/hypebid/go-micro-template/internal/db"
	"github.com/hypebid/go-micro-template/internal/rpc/pb"
)

type Server struct {
	pb.UnsafeServiceNameServer
	Config *config.Config
}

func (s *Server) HealthCheck(ctx context.Context, req *pb.HealthRequest) (*pb.HealthStatus, error) {
	// Get TransactionId from ctx
	tId := ctx.Value(grpc_reqId.TransactionIdMarker("transaction_id_ctx_marker"))
	s.Config.Log.Printf("TransactionId: %v", tId)
	s.Config.Log.Printf("Received: %v", req.GetMessage())

	// ping db
	dbOnline := false
	ping := db.PingDB(s.Config)
	if ping == nil {
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
