package sqlite

import (
	"fmt"
	"mammon/internal/object"
)

func (r *Repository) ReadTasks() ([]object.Task, error) {
	const op = op + ".ReadTask()"

	const query = `
		SELECT name, description, status 
		FROM task;
	`

	var tasks []object.Task
	rows, err := r.db.Query(query)
	if err != nil {
		return tasks, fmt.Errorf("%s | reading tasks failed: %w", op, err)
	}

	for rows.Next() {
		var task object.Task
		err := rows.Scan(&task.Name, &task.Description, &task.Status)
		if err != nil {
			return tasks, fmt.Errorf("%s | scanning tasks failed: %w", op, err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

//func (r *Repository) ReadTask() (object.Task, error) {
//	const op = op + ".ReadTask()"
//
//	const query = `
//		SELECT name, description, status
//		FROM task;
//	`
//
//	var task object.Task
//	err := r.db.QueryRow(query).Scan(&task.Name, &task.Description, &task.Status)
//	if err != nil {
//		return task, fmt.Errorf("%s | reading tasks failed: %w", op, err)
//	}
//
//	return task, nil
//}
