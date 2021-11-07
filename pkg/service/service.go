package service

import (
	"github.com/kagorbunov/todo"
	"github.com/kagorbunov/todo/pkg/reposotory"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *reposotory.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
