package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/kyfk/golang-todo-list-sample/pkg/ent"
)

type ListTodoParams struct {
	Offset int32
	Limit  int32
}

type UpdateTodoParams struct {
	ID    uuid.UUID
	Title string
	Desc  string
}

type CreateTodoParams struct {
	Title string
	Desc  string
}

type Todo interface {
	List(context.Context, ListTodoParams) ([]*ent.Todo, error)
	Get(context.Context, uuid.UUID) (*ent.Todo, error)
	Update(context.Context, UpdateTodoParams) error
	Create(context.Context, CreateTodoParams) error
	Delete(context.Context, uuid.UUID) error
}
