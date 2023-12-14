package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type TodoGW struct {
	once sync.Once

	ServerAddr  string `envconfig:"SERVER_ADDR" required:"true"`
	GatewayAddr string `envconfig:"GATEWAY_ADDR" required:"true"`
	LogLevel    string `envconfig:"LOG_LEVEL" default:"debug"`
}

var (
	todoGW *TodoGW
)

func loadTodoGW() {
	err := envconfig.Process("", todoGW)
	if err != nil {
		panic(err)
	}
}

func GetTodoGW() *TodoGW {
	todoGW.once.Do(loadTodoGW)
	new := *todoGW
	return &new
}
