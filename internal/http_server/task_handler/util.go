package task_handler

import (
	"encoding/json"
	"net/http"
)

const op = "mammon.internal.http-server.task_handler"

func writeJSON(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		// безсмысленно менять status
	}
}
