package server

import (
	"context"

	"github.com/kyfk/golang-todo-list-sample/pkg/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type Health struct {
	*health.Server
	mysql mysql.Driver
}

func NewHealth(mysql mysql.Driver) *Health {
	return &Health{health.NewServer(), mysql}
}

func (h *Health) AuthFuncOverride(ctx context.Context, fullpath string) (context.Context, error) {
	return ctx, nil
}

func (h *Health) Check(ctx context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	if h.mysql != nil {
		if err := h.mysql.Ping(ctx); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return h.Server.Check(ctx, in)
}
