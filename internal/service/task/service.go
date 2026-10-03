package task

import "log/slog"

type Repository interface {
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
