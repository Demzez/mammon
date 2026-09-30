package handler

import (
	"log/slog"
	"net/http"
)

type TaskGetter interface {
	GetTask() string
}

func NewTaskGetter(log *slog.Logger, getter TaskGetter) http.HandlerFunc {
	const op = "mammon.internal.http-server.handler.NewTaskGetter"
	log = log.With(slog.String("op", op))

	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("TaskGetter successfully")
		w.Write([]byte(getter.GetTask()))
		return
	}
}
