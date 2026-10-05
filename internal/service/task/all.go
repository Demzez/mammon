package task

import (
	"log/slog"
	"mammon/internal/object"
)

func (s *Service) GetTasks() ([]object.Task, error) {
	const op = op + ".GetTasks"
	log := s.log.With("op", op)

	tasks, err := s.rep.ReadTasks()
	if err != nil {
		log.Error("task getting failed", slog.String("error", err.Error()))
		return tasks, err
	}

	return tasks, nil
}
