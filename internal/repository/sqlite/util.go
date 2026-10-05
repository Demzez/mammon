package sqlite

import "fmt"

func (r *Repository) InitSchema() error {
	const op = op + ".InitSchema"

	const query = `
	CREATE TABLE IF NOT EXISTS task (
		task_id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		status TEXT NOT NULL
	);
	`

	if _, err := r.db.Exec(query); err != nil {
		return fmt.Errorf("%s | schema init failed: %w", op, err)
	}

	return nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}
