package main

import (
	"fmt"
	"log/slog"
	"mammon/internal/config"
	"mammon/internal/lib/logger/slog/slogpretty"
	"mammon/internal/repository/sqlite"
	"mammon/internal/service/task"
	"net/http"
	"os"
	"sync"
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
	repository, err := sqlite.NewRepository(cfg)
	if err != nil {
		log.Error("database init failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer repository.Close()
	log.Info("database is loaded")

	//init services
	services := Services{
		task: task.NewService(logger, repository),
	}
	log.Info("services is created")

	// start http server
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	router := createRouter(logger, services)

	wg.Add(1)
	go func() {
		http.ListenAndServe(address, router)
		wg.Done()
	}()
	log.Info("server is started on " + address)

	//main goroutine wait until server work
	wg.Wait()
}
