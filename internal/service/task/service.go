package task

import (
	"log/slog"
	"mammon/internal/object"
)

const op = "mammon.internal.service.task"

type Repository interface {
	ReadTasks() ([]object.Task, error)
}

type Service struct {
	log *slog.Logger
	rep Repository
}

func NewService(log *slog.Logger, rep Repository) *Service {
	return &Service{
		log: log,
		rep: rep,
	}
}
