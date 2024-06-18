package repository

import (
	newtodo "exapmle/todo_app"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user newtodo.User) (int, error)
	GetUser(username, password string) (newtodo.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
