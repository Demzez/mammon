package task_handler

import (
	"encoding/json"
	"log/slog"
	"mammon/internal/object"
	"net/http"
)

type TaskAdder interface {
	AddTask(task object.Task) (int64, error)
}

func AddTask(l *slog.Logger, adder TaskAdder) http.HandlerFunc {
	const op = op + ".AddTask()"
	log := l.With("op", op)

	return func(w http.ResponseWriter, r *http.Request) {
		task := object.Task{}

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			log.Error("failed to unmarshal json", slog.String("error", err.Error()))
		}

		taskId, err := adder.AddTask(task)
		if err != nil {
			log.Error("fialed to add task", slog.String("error", err.Error()))
		}

		_ = taskId

		w.WriteHeader(http.StatusCreated)
	}
}
