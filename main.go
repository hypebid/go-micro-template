package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/hypebid/go-micro-template/internal/config"
	"github.com/hypebid/go-micro-template/internal/drpc"
	"github.com/hypebid/go-micro-template/internal/drpc/pb"
	"github.com/hypebid/go-micro-template/internal/rest/controllers"
	"golang.org/x/sync/errgroup"
	"storj.io/drpc/drpchttp"
	"storj.io/drpc/drpcmigrate"
	"storj.io/drpc/drpcmux"
	"storj.io/drpc/drpcserver"
)

func main() {
	err := start(context.Background())
	if err != nil {
		panic(err)
	}
}

func routes(c *config.Config) *chi.Mux {
	r := chi.NewRouter()

	// TODO: set up logger

	// add middleware to router
	r.Use(render.SetContentType(render.ContentTypeJSON),
		middleware.RequestID,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/heartbeat", controllers.HeartbeatRoutes(c))
	})

	return r
}

func start(ctx context.Context) error {
	// create config
	// c := config.New()

	// create a RPC server
	cookieMonster := &drpc.CookieMonsterServer{}

	// create a dRPC RPC mux
	drpcMux := drpcmux.New()

	// register the proto-spcecific methods on the mux
	err := pb.DRPCRegisterCookieMonster(drpcMux, cookieMonster)
	if err != nil {
		return err
	}

	// listen on a tcp socket
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	// create a listen mux that evalutes enough bytes to recognize the dRPC header
	lisMux := drpcmigrate.NewListenMux(lis, len(drpcmigrate.DRPCHeader))

	// creating group to run different protocol servers in parallel.
	var group errgroup.Group

	// drpc handling
	group.Go(func() error {
		// create a dRPC server
		drpcServer := drpcserver.New(drpcMux)

		// grab the listen mux route for the dRPC header
		drpcLis := lisMux.Route(drpcmigrate.DRPCHeader)

		//run the server
		log.Println("dRPC server running...")
		return drpcServer.Serve(ctx, drpcLis)
	})

	// http handling for dRPC methods
	group.Go(func() error {
		// create an http server
		// router := routes(c, drpcMux)
		httpServer := http.Server{Handler: drpchttp.New(drpcMux)}

		// run the server
		log.Println("http server running...")
		return httpServer.Serve(lisMux.Default())
	})

	// run the listen mux
	group.Go(func() error {
		return lisMux.Run(ctx)
	})

	// wait
	log.Println("service running on port :8080")
	return group.Wait()
}
