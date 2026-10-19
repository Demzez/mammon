package task_handler

import (
	"net/http"
)

const op = "mammon.internal.http_server.task_handler"

func writeJSON(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err := w.Write(data)
	if err != nil {
		// менять статус поздно
	}
}

const HttpError = "something went wrong"
