package main

import (
	"log/slog"
	"mammon/internal/http-server/task_handler"
	"mammon/internal/service/task"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Services struct {
	task *task.Service
}

func createRouter(log *slog.Logger, services Services) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.URLFormat)

	r.Get("/tasks", task_handler.GetTasks(log, services.task))

	return r
}
