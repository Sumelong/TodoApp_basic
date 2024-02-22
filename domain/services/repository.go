package services

import (
	"TodoApp_basic/domain/entity"
	"TodoApp_basic/domain/query"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(task *entity.Task) (string, error) {

	tx, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(query.Add)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	result, err := stmt.Exec(task.Id, task.CreatedAt, task.UpdatedAt, task.Item, task.Done, task.DoneAt)
	if err != nil {
		log.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return "", err
	}

	return task.Id, nil
}

func (r *Repository) Update(task *entity.Task) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(query.Update)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	result, err := stmt.Exec(&task.UpdatedAt, &task.Item, &task.Done, &task.DoneAt, &task.Id)
	if err != nil {
		log.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return "", err
	}

	return task.Id, nil

}

func (r *Repository) FindAll() ([]entity.Task, error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(query.FindAll)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to hold tasks
	var tasks []entity.Task

	// Iterate over the result rows
	for rows.Next() {
		// Create a new Task instance
		var task entity.Task
		// Scan the columns of the current row into the fields of the Task struct
		if err = rows.Scan(&task.Id, &task.CreatedAt, &task.UpdatedAt, &task.Item, &task.Done, &task.DoneAt); err != nil {
			return nil, err
		}
		// Append the task to the slice
		tasks = append(tasks, task)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil

}

func (r *Repository) FindBy(where *entity.Task) (task entity.Task, err error) {

	// Initialize a slice to hold tasks
	var NewTask entity.Task

	tx, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(query.FindBy)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	row := stmt.QueryRow(&where.Id)
	if err != nil {
		log.Fatal(err)
		return task, err
	}

	// Scan the columns of the current row into the fields of the Task struct
	err = row.Scan(&task.Id, &task.CreatedAt, &task.UpdatedAt, &task.Item, &task.Done, &task.DoneAt)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			// Handle case where no rows were returned
			return NewTask, fmt.Errorf("no task found with ID %s", task.Id)
		}
		// Handle other errors
		return NewTask, err
	}

	return task, nil

}

func (r *Repository) Remove(task *entity.Task) (string, error) {

	tx, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(query.Remove)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	result, err := stmt.Exec(task.Id)
	if err != nil {
		log.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return "", err
	}

	return task.Id, nil
}
