package main

import (
	"log/slog"
	"mammon/internal/config"
	"mammon/internal/http-server/handler"
	"mammon/internal/lib/logger/slog/slogpretty"
	"mammon/internal/repository/sqlite"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	const op = "mammon.cmd.mammon.main"

	//read config
	cfg := config.MustLoad()

	// init logger
	logger := slogpretty.SetupPrettyLogger()
	log := logger.With(slog.String("op", op))
	log.Info("config is read && logger is loaded")

	// init database
	repository := sqlite.NewSqliteRepository(cfg, logger)
	defer repository.Close()
	log.Info("database is loaded")

	// start http server
	r := chi.NewRouter()
	r.Use(middleware.URLFormat)
	r.Get("/", handler.NewTextSender(logger))
	http.ListenAndServe(":5555", r)
}
