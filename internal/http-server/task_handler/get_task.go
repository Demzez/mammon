package task_handler

import (
	"encoding/json"
	"log/slog"
	"mammon/internal/object"
	"net/http"
)

type TaskGetter interface {
	GetTasks() ([]object.Task, error)
}

func GetTasks(l *slog.Logger, getter TaskGetter) http.HandlerFunc {
	const op = op + ".GetTasks()"
	log := l.With(slog.String("op", op))

	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := getter.GetTasks()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(tasks)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(data)
		log.Info("tasks get successfully")
		return
	}
}
