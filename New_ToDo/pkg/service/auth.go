package service

import (
	"crypto/sha1"
	newtodo "exapmle/todo_app"
	"exapmle/todo_app/pkg/repository"
	"fmt"
)

const solt = "dsaddjui1j2jk3e2193mjaklsd"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user newtodo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(solt)))
}
