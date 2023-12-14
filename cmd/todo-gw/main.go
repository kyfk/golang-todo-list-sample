package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kyfk/golang-todo-list-sample/config"
	gw "github.com/kyfk/golang-todo-list-sample/gateway/todo-list/server/v1"
	"github.com/kyfk/golang-todo-list-sample/pkg/build"
	"github.com/kyfk/golang-todo-list-sample/pkg/grpcgw/middleware"
	"github.com/kyfk/golang-todo-list-sample/pkg/log"
	"github.com/kyfk/golang-todo-list-sample/pkg/mysql"
	"github.com/oklog/run"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	cfg := config.GetTodoGW()

	lv, err := zapcore.ParseLevel(cfg.LogLevel)
	if err != nil {
		panic(err)
	}

	sc := log.ServiceContext{
		Service:   "todo-gw",
		Version:   build.Version,
		GitCommit: build.GitCommit,
		BuildDate: build.BuildDate,
	}
	logger, err := log.NewLogger(lv, sc)
	if err != nil {
		panic(err)
	}

	if err := mysql.SetErrLogger(logger); err != nil {
		logger.Fatal("failed to set mysql error logger", zap.Error(err))
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var g run.Group

	{
		term := make(chan os.Signal, 1)
		signal.Notify(term, syscall.SIGTERM)
		g.Add(func() error {
			select {
			case sig := <-term:
				logger.Info("signal received", zap.String("signal", sig.String()))
				return errors.New(sig.String())
			case <-ctx.Done():
				return errors.New("context canceled")
			}
		}, func(err error) {})
	}

	{
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		conn, err := grpc.Dial(cfg.ServerAddr, opts...)
		if err != nil {
			logger.Fatal("failed to dial endpoint", zap.Error(err))
		}
		logger.Debug("start gateway http server at", zap.String("address", cfg.GatewayAddr))
		mux := runtime.NewServeMux(
			runtime.WithHealthzEndpoint(healthpb.NewHealthClient(conn)),
			runtime.WithErrorHandler(middleware.ErrorHandler),
		)
		if err := gw.RegisterTodoHandler(ctx, mux, conn); err != nil {
			logger.Fatal("failed to register recognition handler", zap.Error(err))
		}

		server := &http.Server{Addr: cfg.GatewayAddr, Handler: mux}
		g.Add(func() error {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Error("failed to serve gateway", zap.Error(err))
				return err
			}
			return nil
		}, func(err error) {
			logger.Debug(err.Error())
			if err := server.Shutdown(ctx); err != nil {
				logger.Error("failed to shutdown gateway", zap.Error(err))
			}
		})
	}

	if err := g.Run(); err != nil {
		os.Exit(1)
	}
}
