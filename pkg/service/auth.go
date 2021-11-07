package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/kagorbunov/todo"
	"github.com/kagorbunov/todo/pkg/reposotory"
)

const salt = "fcuv;hn[shgdnvwsnregw"

type AuthService struct {
	repo reposotory.Authorization
}

func NewAuthService(repo reposotory.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
