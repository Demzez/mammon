package sqlite

import (
	"database/sql"
	"fmt"
	"mammon/internal/config"

	_ "github.com/mattn/go-sqlite3"
)

const op = "mammon.internal.reposiory.sqlite"

type Repository struct {
	db *sql.DB
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	const op = op + ".NewRepository"

	db, err := sql.Open("sqlite3", cfg.DataBase.Path)
	if err != nil {
		return nil, fmt.Errorf("%s | database open failed: %w", op, err)
	}

	repository := &Repository{
		db: db,
	}

	if err := repository.InitSchema(); err != nil {
		db.Close()
		return nil, err
	}

	return repository, nil
}
