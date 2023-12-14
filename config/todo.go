package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/kyfk/golang-todo-list-sample/pkg/env"
	"github.com/kyfk/golang-todo-list-sample/pkg/mysql"
)

type Todo struct {
	once sync.Once

	Env          env.Env      `envconfig:"ENV" required:"true"`
	MySQL        mysql.Config `envconfig:"MYSQL" required:"true"`
	Port         string       `envconfig:"PORT" required:"true"`
	MetricsPort  string       `envconfig:"METRICS_PORT" required:"true"`
	LogLevel     string       `envconfig:"LOG_LEVEL" default:"debug"`
	GrpcLogLevel string       `envconfig:"GRPC_LOG_LEVEL" default:"warn"`
}

var (
	todo *Todo
)

func loadTodo() {
	err := envconfig.Process("", todo)
	if err != nil {
		panic(err)
	}
}

func GetTodo() *Todo {
	todo.once.Do(loadTodo)
	new := *todo
	return &new
}
