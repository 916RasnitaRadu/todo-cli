package repository

import "github.com/916RasnitaRadu/todo-cli/types"

type Repository interface {
	GetTasks() ([]types.Task, error)
	Create(task types.Task) error
	Delete(id int) error
	ChangeStatus(id int) error
}
