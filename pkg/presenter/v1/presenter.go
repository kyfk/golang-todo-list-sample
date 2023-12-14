package v1

import (
	"time"

	"github.com/kyfk/golang-todo-list-sample/pkg/ent"
	todoucv1 "github.com/kyfk/golang-todo-list-sample/pkg/usecase/todo/v1"
	modelv1 "github.com/kyfk/golang-todo-list-sample/server/interface/model/v1"
	v1 "github.com/kyfk/golang-todo-list-sample/server/interface/todo/v1"
)

type Presenter struct {
}

func New() *Presenter {
	return &Presenter{}
}

func (p *Presenter) Index(out *todoucv1.IndexOutput) *v1.IndexResponse {
	return &v1.IndexResponse{
		Todos: p.Todos(out.Todos),
	}
}

func (p *Presenter) Get(out *todoucv1.GetOutput) *v1.GetResponse {
	return &v1.GetResponse{
		Todo: p.Todo(out.Todo),
	}
}

func (p *Presenter) Todos(todos []*ent.Todo) (r []*modelv1.Todo) {
	r = make([]*modelv1.Todo, len(todos))
	for i := range todos {
		r[i] = p.Todo(todos[i])
	}
	return
}

func (p *Presenter) Todo(todo *ent.Todo) *modelv1.Todo {
	return &modelv1.Todo{
		Id:        todo.ID.String(),
		Title:     todo.Title,
		Desc:      todo.Desc,
		CreatedAt: todo.CreatedAt.Format(time.RFC3339),
		UpdatedAt: todo.UpdatedAt.Format(time.RFC3339),
	}
}
