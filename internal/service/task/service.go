package task

import (
	"mammon/internal/object"
)

const op = "mammon.internal.service.task"

type Repository interface {
	CreateTask(task object.Task) (int64, error)
	ReadTasks() ([]object.Task, error)
}

type Service struct {
	rep Repository
}

func NewService(rep Repository) *Service {
	return &Service{
		rep: rep,
	}
}
