package task

import (
	"database/sql"
	"fmt"

	"github.com/victorluisca/go-todo-app/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetAllTasks() ([]*types.Task, error) {
	rows, err := s.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	var tasks []*types.Task

	for rows.Next() {
		task, err := scanRowsIntoTask(rows)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *Store) CreateTask(task types.Task) error {
	_, err := s.db.Exec("INSERT INTO tasks (title, priority) VALUES (?, ?)", task.Title, task.Priority)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetTaskByID(id int) (*types.Task, error) {
	rows, err := s.db.Query("SELECT * FROM tasks WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("task with ID %d not found", id)
	}

	task, err := scanRowsIntoTask(rows)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *Store) UpdateTask(task types.Task) error {
	_, err := s.db.Exec("UPDATE tasks SET title = ?, priority = ? WHERE id = ?", task.Title, task.Priority, task.ID)
	return err
}

func (s *Store) DeleteTask(id int) error {
	_, err := s.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

func scanRowsIntoTask(rows *sql.Rows) (*types.Task, error) {
	task := new(types.Task)

	err := rows.Scan(
		&task.ID,
		&task.Title,
		&task.Priority,
		&task.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return task, err
}
