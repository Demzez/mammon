package task

import "log/slog"

type Service struct {
	log *slog.Logger
	//TODO: repository interface
}

func NewService(log *slog.Logger) *Service {
	return &Service{
		log: log,
	}
}
