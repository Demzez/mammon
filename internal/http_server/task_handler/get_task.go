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
			http.Error(w, "something went wrong", http.StatusNoContent)

			log.Error("failed to getting tasks from service", slog.String("error", err.Error()))
			return
		}

		data, err := json.Marshal(tasks)
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)

			log.Error("failed to marshal json", slog.String("error", err.Error()))
			return
		}

		writeJSON(w, http.StatusOK, data)
		log.Info("successfully get all tasks")
		return
	}
}
