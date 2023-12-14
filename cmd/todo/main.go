package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/kyfk/golang-todo-list-sample/config"
	"github.com/kyfk/golang-todo-list-sample/pkg/build"

	_ "github.com/kyfk/golang-todo-list-sample/pkg/ent/runtime"
	"github.com/kyfk/golang-todo-list-sample/pkg/log"
	"github.com/kyfk/golang-todo-list-sample/pkg/mysql"
	"github.com/kyfk/golang-todo-list-sample/pkg/server"
	todov1pb "github.com/kyfk/golang-todo-list-sample/server/interface/todo/v1"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	cfg := config.GetTodo()

	lv, err := zapcore.ParseLevel(cfg.LogLevel)
	if err != nil {
		panic(err)
	}

	sc := log.ServiceContext{
		Service:   "todo-app",
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mysqlDriver, err := mysql.NewDriver(config.GetTodo().MySQL)
	if err != nil {
		logger.Fatal("mysql initialization failed")
	}
	defer mysqlDriver.Close()

	glv, err := zapcore.ParseLevel(cfg.LogLevel)
	if err != nil {
		logger.Fatal("grpc logger level parse error")
	}
	grpcLogger, err := log.NewGRPCLogger(glv, sc)
	if err != nil {
		logger.Fatal("grpc logger initialization failed")
	}
	grpclog.SetLoggerV2(zapgrpc.NewLogger(grpcLogger.With(zap.String("group", "grpc"))))

	appServer, err := initialize(
		ctx,
		mysqlDriver,
		logger,
	)
	if err != nil {
		logger.Fatal("failed to initialize manager", zap.Error(err))
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(RecoveryHandler)),
			grpc_validator.UnaryServerInterceptor(),
		)),
	)

	if cfg.Env.IsLocal() || cfg.Env.IsDev() || cfg.Env.IsQA() {
		logger.Debug("reflection enabled")
		reflection.Register(grpcServer)
	}

	todov1pb.RegisterTodoServer(grpcServer, appServer)
	healthpb.RegisterHealthServer(grpcServer, server.NewHealth(mysqlDriver))
	grpc_prometheus.Register(grpcServer)
	grpc_prometheus.EnableHandlingTimeHistogram()

	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		logger.Fatal("unexpected error", zap.Error(err))
	}

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
		logger.Debug("start grpc server at", zap.String("address", ":"+cfg.Port))
		g.Add(func() error {
			if err := grpcServer.Serve(lis); err != nil {
				logger.Error("failed to serve", zap.Error(err))
				return err
			}
			return nil
		}, func(error) {
			grpcServer.GracefulStop()
		})
	}

	{
		logger.Debug("start metrics http server at", zap.String("address", ":"+cfg.MetricsPort))
		mux := http.DefaultServeMux
		mux.Handle("/metrics", promhttp.Handler())
		server := &http.Server{Addr: ":" + cfg.MetricsPort, Handler: mux}
		g.Add(func() error {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Error("failed to serve metrics", zap.Error(err))
				return err
			}
			return nil
		}, func(err error) {
			logger.Debug(err.Error())
			if err := server.Shutdown(ctx); err != nil {
				logger.Error("failed to shutdown metrics", zap.Error(err))
			}
		})
	}

	if err := g.Run(); err != nil {
		os.Exit(1)
	}
}

func RecoveryHandler(ctx context.Context, p interface{}) (err error) {
	ctxzap.Extract(ctx).Error("panicked", zap.Any("p", p))
	return status.Errorf(codes.Unknown, "panic triggered: %v", p)
}

func OmitHealth(fullMethodName string, _ error) bool {
	return fullMethodName != "/grpc.health.v1.Health/Check"
}
