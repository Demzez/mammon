package handler

import (
	"log/slog"
	"net/http"
)

func NewTextSender(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello?, mammon"))
		return
	}
}
