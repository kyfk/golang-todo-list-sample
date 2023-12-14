package server

import (
	context "context"

	"github.com/google/uuid"
	p1 "github.com/kyfk/golang-todo-list-sample/pkg/presenter/v1"
	todoucv1 "github.com/kyfk/golang-todo-list-sample/pkg/usecase/todo/v1"
	v1 "github.com/kyfk/golang-todo-list-sample/server/interface/todo/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type TodoV1Server struct {
	v1.UnimplementedTodoServer

	p1 *p1.Presenter

	todo *todoucv1.Usecase
}

func NewTodoV1Server(p1 *p1.Presenter, todo *todoucv1.Usecase) *TodoV1Server {
	return &TodoV1Server{
		p1:   p1,
		todo: todo,
	}
}

func (s *TodoV1Server) Index(ctx context.Context, req *v1.IndexRequest) (*v1.IndexResponse, error) {
	out, err := s.todo.Index(ctx, &todoucv1.IndexInput{
		Offset: req.Offset,
		Limit:  req.Limit,
	})
	if err != nil {
		return nil, err
	}
	return s.p1.Index(out), nil
}

func (s *TodoV1Server) Get(ctx context.Context, req *v1.GetRequest) (*v1.GetResponse, error) {
	id, _ := uuid.Parse(req.Id)
	out, err := s.todo.Get(ctx, &todoucv1.GetInput{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	return s.p1.Get(out), nil
}

func (s *TodoV1Server) Update(ctx context.Context, req *v1.UpdateRequest) (*emptypb.Empty, error) {
	id, _ := uuid.Parse(req.Id)
	err := s.todo.Update(ctx, &todoucv1.UpdateInput{
		ID:    id,
		Title: req.Title,
		Desc:  req.Desc,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *TodoV1Server) Create(ctx context.Context, req *v1.CreateRequest) (*emptypb.Empty, error) {
	err := s.todo.Create(ctx, &todoucv1.CreateInput{
		Title: req.Title,
		Desc:  req.Desc,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *TodoV1Server) Delete(ctx context.Context, req *v1.DeleteRequest) (*emptypb.Empty, error) {
	id, _ := uuid.Parse(req.Id)
	err := s.todo.Delete(ctx, &todoucv1.DeleteInput{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
