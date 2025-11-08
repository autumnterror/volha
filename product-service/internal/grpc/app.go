package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"productService/config"
	"productService/internal/pkg/psql"
	"productService/internal/utils/format"
)

type App struct {
	gRPCServer *grpc.Server
	cfg        *config.Config
	API        psql.Repository
}

// New construct new App structure
func New(
	cfg *config.Config,
	API psql.Repository,
) *App {
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 0,
		}),
	)
	Register(s, API)
	return &App{
		gRPCServer: s,
		cfg:        cfg,
		API:        API,
	}
}

// MustRun running gRPC server and panic if error
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpc.App"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.cfg.Port))
	if err != nil {
		return format.Error(op, err)
	}
	log.Println(op, "grpc server is running", a.cfg.Port)

	if err := a.gRPCServer.Serve(l); err != nil {
		return format.Error(op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpc.Stop"
	a.gRPCServer.GracefulStop()
	log.Println(op, "grpc server is stop", a.cfg.Port)
}
