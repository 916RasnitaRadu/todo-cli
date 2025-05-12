package service

import (
	"github.com/916RasnitaRadu/todo-cli/repository"
	"github.com/916RasnitaRadu/todo-cli/types"
)

type Service struct {
	repo repository.Repository
}

func (s *Service) GetTasks() ([]types.Task, error) {
	tasks, err := s.repo.GetTasks()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *Service) Create(task types.Task) error {
	if err := s.repo.Create(task); err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (s *Service) ChangeStatus(id int) error {
	if err := s.repo.ChangeStatus(id); err != nil {
		return err
	}

	return nil
}
