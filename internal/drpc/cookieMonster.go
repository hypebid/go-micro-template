package drpc

import (
	"context"

	"github.com/hypebid/go-micro-template/internal/drpc/pb"
)

type CookieMonsterServer struct {
	pb.DRPCCookieMonsterUnimplementedServer
}

// EatCookie : dRPC method that turns a cookie into crumbs.
func (s *CookieMonsterServer) EatCookie(ctx context.Context, cookie *pb.Cookie) (*pb.Crumbs, error) {
	return &pb.Crumbs{
		Cookie: cookie,
	}, nil
}
