package rpc

import (
	"context"
	"runtime"

	"github.com/hypebid/go-kit/grpc/middleware"
	"github.com/hypebid/go-micro-template/internal/config"
	"github.com/hypebid/go-micro-template/internal/db"
	"github.com/hypebid/go-micro-template/internal/rpc/pb"
	"github.com/sirupsen/logrus"
)

type Server struct {
	pb.UnsafeServiceNameServer
	Config *config.Config
}

func (s *Server) HealthCheck(ctx context.Context, req *pb.HealthRequest) (*pb.HealthStatus, error) {
	// Build logger with TransactionId
	tId := ctx.Value(middleware.Grpc_ReqId_Marker)
	pc, _, _, _ := runtime.Caller(0)
	logger := s.Config.Log.WithFields(logrus.Fields{"transaction-id": tId, "method": runtime.FuncForPC(pc).Name()})

	logger.Printf("received: %v", req.GetMessage())

	// ping db
	dbOnline := false
	ping := db.PingDB(s.Config)
	if ping == nil {
		dbOnline = true
	}

	return &pb.HealthStatus{
		TransactionId:  tId.(string),
		ServiceName:    s.Config.Constants.ServiceName,
		ReleaseDate:    s.Config.Constants.ReleaseDate,
		ReleaseSlug:    s.Config.Constants.ReleaseSlug,
		ReleaseVersion: s.Config.Constants.ReleaseVersion,
		DatabaseOnline: dbOnline,
		Message:        req.GetMessage(),
	}, nil
}
