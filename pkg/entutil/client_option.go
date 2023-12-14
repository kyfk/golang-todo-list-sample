package entutil

import (
	"github.com/kyfk/golang-todo-list-sample/pkg/ent"
	"github.com/kyfk/golang-todo-list-sample/pkg/env"
	"github.com/kyfk/golang-todo-list-sample/pkg/mysql"
)

func NewOptions(env env.Env, mysql mysql.Driver) []ent.Option {
	options := []ent.Option{
		ent.Driver(mysql),
	}
	if env.IsQA() || env.IsDev() || env.IsLocal() {
		options = append(options, ent.Debug())
	}
	return options
}
