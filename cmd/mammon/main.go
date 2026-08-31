package main

import (
	"fmt"
	"log/slog"
	"mammon/internal/config"
	"mammon/internal/http-server/handler"
	"mammon/internal/lib/logger/slog/slogpretty"
	"mammon/internal/repository/sqlite"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	const op = "mammon.cmd.mammon.main"

	wg := &sync.WaitGroup{}

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
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	router := createRouter(logger)

	wg.Add(1)
	go func() {
		http.ListenAndServe(address, router)
		wg.Done()
	}()

	log.Info("server is started on " + address)

	//main goroutine wait until server work
	wg.Wait()
}

func createRouter(log *slog.Logger) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.URLFormat)

	r.Get("/", handler.NewTextSender(log))

	return r
}
