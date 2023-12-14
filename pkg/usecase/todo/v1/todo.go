package v1

import (
	"context"

	"github.com/google/uuid"
	"github.com/kyfk/golang-todo-list-sample/pkg/ent"
	"github.com/kyfk/golang-todo-list-sample/pkg/usecase/repo"
	"github.com/pkg/errors"
)

type Usecase struct {
	todoRepo repo.Todo
}

func New(todoRepo repo.Todo) *Usecase {
	return &Usecase{todoRepo}
}

type IndexInput struct {
	Offset int32
	Limit  int32
}

type IndexOutput struct {
	Todos []*ent.Todo
}

func (u *Usecase) Index(ctx context.Context, input *IndexInput) (*IndexOutput, error) {
	todos, err := u.todoRepo.List(ctx, repo.ListTodoParams{
		Limit:  input.Limit,
		Offset: input.Offset,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &IndexOutput{
		Todos: todos,
	}, nil
}

type GetInput struct {
	ID uuid.UUID
}

type GetOutput struct {
	Todo *ent.Todo
}

func (u *Usecase) Get(ctx context.Context, input *GetInput) (*GetOutput, error) {
	todo, err := u.todoRepo.Get(ctx, input.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &GetOutput{
		Todo: todo,
	}, nil
}

type UpdateInput struct {
	ID    uuid.UUID
	Title string
	Desc  string
}

type UpdateOutput struct {
	Todo *ent.Todo
}

func (u *Usecase) Update(ctx context.Context, input *UpdateInput) error {
	err := u.todoRepo.Update(ctx, repo.UpdateTodoParams{
		ID:    input.ID,
		Title: input.Title,
		Desc:  input.Desc,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

type CreateInput struct {
	Title string
	Desc  string
}

func (u *Usecase) Create(ctx context.Context, input *CreateInput) error {
	err := u.todoRepo.Create(ctx, repo.CreateTodoParams{
		Title: input.Title,
		Desc:  input.Desc,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

type DeleteInput struct {
	ID uuid.UUID
}

func (u *Usecase) Delete(ctx context.Context, input *DeleteInput) error {
	err := u.todoRepo.Delete(ctx, input.ID)
	return errors.WithStack(err)
}
