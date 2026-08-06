package sqlite

import (
	"database/sql"
	"log/slog"
	"mammon/internal/config"

	_ "github.com/mattn/go-sqlite3"
)

const op = "mammon.internal.repository.sqlite"

func NewSqliteRepository(cfg *config.Config, log *slog.Logger) *sql.DB {
	log = log.With(
		slog.String("op", op+"NewSqliteRepository"),
	)

	db, err := sql.Open("sqlite3", cfg.DataBase.Path)
	if err != nil {
		log.Error("failed to open database", slog.String("error", err.Error()))
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS task (
    	task_id INTEGER PRIMARY KEY,
    	name TEXT NOT NULL,
    	descritption TEXT,
        status TEXT NOT NULL)
        `)
	if err != nil {
		log.Error("failed to create table", "error:", err)
		db.Close()
	}

	return db
}

//func
