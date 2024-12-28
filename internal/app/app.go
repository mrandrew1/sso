package app

import (
	"log/slog"
	"time"

	grpcapp "sso/internal/grpcapp"
)

type App struct {
	GRPCSever *grpcapp.Server
}

func New(log *slog.Logger, grpcPort int, StoragePath string, tokerTTl time.Duration) *App {

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSever: grpcApp,
	}
}
