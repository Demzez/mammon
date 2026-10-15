package task

import (
	"fmt"
	"mammon/internal/object"
)

func (s *Service) AddTask(task object.Task) (int64, error) {
	const op = op + ".AddTask()"

	taskId, err := s.rep.CreateTask(task)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return taskId, nil
}

func (s *Service) GetTasks() ([]object.Task, error) {
	const op = op + ".GetTasks()"

	tasks, err := s.rep.ReadTasks()
	if err != nil {
		return tasks, fmt.Errorf("%s : %w", op, err)
	}

	return tasks, nil
}
