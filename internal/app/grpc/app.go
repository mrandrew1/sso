package grpcapp

import (
	"fmt"
	"log/slog"
	"net"
	authgrpc "sso/internal/grpc/auth"

	"google.golang.org/grpc"
)

type App struct {
	log       *slog.Logger
	gRPCSever *grpc.Server
	port      int
}

func New(log *slog.Logger, port int) *App {
	gRPCSever := grpc.NewServer()
	authgrpc.Register(gRPCSever)

	return &App{
		log:       log,
		gRPCSever: gRPCSever,
		port:      port,
	}
}
func (a *App) MustRun() {
	if err := a.Run; err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("running gRPC server", slog.String("address", l.Addr().String()))

	if err := a.gRPCSever.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).
		Info("stopping grpc server", slog.Int("port", a.port))

	a.gRPCSever.GracefulStop()
}
