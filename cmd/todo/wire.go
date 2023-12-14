//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"github.com/kyfk/golang-todo-list-sample/config"
	"github.com/kyfk/golang-todo-list-sample/pkg/ent"
	"github.com/kyfk/golang-todo-list-sample/pkg/entutil"
	"github.com/kyfk/golang-todo-list-sample/pkg/mysql"
	p1 "github.com/kyfk/golang-todo-list-sample/pkg/presenter/v1"
	repoimpl "github.com/kyfk/golang-todo-list-sample/pkg/repo"
	"github.com/kyfk/golang-todo-list-sample/pkg/usecase/repo"
	todoucv1 "github.com/kyfk/golang-todo-list-sample/pkg/usecase/todo/v1"
	"github.com/kyfk/golang-todo-list-sample/server"
	"go.uber.org/zap"
)

func initialize(
	ctx context.Context,
	driver mysql.Driver,
	logger *zap.Logger,
) (*server.TodoV1Server, error) {
	wire.Build(
		wire.FieldsOf(new(*config.Todo), "Env"),
		config.GetTodo,
		ent.NewClient,
		entutil.NewOptions,
		repoimpl.NewTodo,
		wire.Bind(new(repo.Todo), new(*repoimpl.Todo)),
		todoucv1.New,
		p1.New,
		server.NewTodoV1Server,
	)
	return &server.TodoV1Server{}, nil
}
