package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/kyfk/golang-todo-list-sample/pkg/ent"
	"github.com/kyfk/golang-todo-list-sample/pkg/usecase/repo"
	"github.com/pkg/errors"
)

type Todo struct {
	client *ent.Client
}

func NewTodo(client *ent.Client) *Todo {
	return &Todo{
		client,
	}
}

func (t *Todo) List(ctx context.Context, params repo.ListTodoParams) ([]*ent.Todo, error) {
	todos, err := t.client.Todo.Query().Limit(int(params.Limit)).Offset(int(params.Offset)).All(ctx)
	return todos, errors.WithStack(err)
}

func (t *Todo) Get(ctx context.Context, id uuid.UUID) (*ent.Todo, error) {
	todo, err := t.client.Todo.Get(ctx, id)
	return todo, errors.WithStack(err)
}

func (t *Todo) Update(ctx context.Context, params repo.UpdateTodoParams) error {
	err := t.client.Todo.
		UpdateOneID(params.ID).
		SetTitle(params.Title).
		SetDesc(params.Desc).
		Exec(ctx)
	return errors.WithStack(err)
}

func (t *Todo) Create(ctx context.Context, params repo.CreateTodoParams) error {
	err := t.client.Todo.
		Create().
		SetID(uuid.New()).
		SetTitle(params.Title).
		SetDesc(params.Desc).
		Exec(ctx)
	return errors.WithStack(err)
}

func (t *Todo) Delete(ctx context.Context, id uuid.UUID) error {
	err := t.client.Todo.
		DeleteOneID(id).
		Exec(ctx)
	return errors.WithStack(err)
}
