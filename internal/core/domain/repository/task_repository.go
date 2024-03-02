package repository

import (
	"TodoApp_basic/internal/core/application/ports/query"
	"TodoApp_basic/internal/core/domain/entity"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type TaskRepository struct {
	db  *sql.DB
	qry query.Query
}

func NewTaskRepository(db *sql.DB, qry query.Query) *TaskRepository {
	return &TaskRepository{
		db:  db,
		qry: qry,
	}
}

func (r *TaskRepository) Create(eTask *entity.Task) (string, error) {

	tx, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(r.qry.Add)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	result, err := stmt.Exec(eTask.Id, eTask.CreatedAt, eTask.UpdatedAt, eTask.Item, eTask.Done, eTask.DoneAt)
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

	return eTask.Id, nil
}

func (r *TaskRepository) Update(eTask *entity.Task) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(r.qry.Update)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	result, err := stmt.Exec(&eTask.UpdatedAt, &eTask.Item, &eTask.Done, &eTask.DoneAt, &eTask.Id)
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

	return eTask.Id, nil

}

func (r *TaskRepository) FindAll() ([]entity.Task, error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(r.qry.FindAll)
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
	var eTasks []entity.Task

	// Iterate over the result rows
	for rows.Next() {
		// Create a new Task instance
		var eTask entity.Task
		// Scan the columns of the current row into the fields of the Task struct
		if err = rows.Scan(&eTask.Id, &eTask.CreatedAt, &eTask.UpdatedAt, &eTask.Item, &eTask.Done, &eTask.DoneAt); err != nil {
			return nil, err
		}
		// Append the taskusecase service to the slice
		eTasks = append(eTasks, eTask)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return eTasks, nil

}

func (r *TaskRepository) FindOne(where *entity.Task) (*entity.Task, error) {

	// Initialize a slice to hold tasks
	var task entity.Task

	tx, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(r.qry.FindOne)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	row := stmt.QueryRow(&where.Id)
	if err != nil {
		log.Fatal(err)
		return &entity.Task{}, err
	}

	// Scan the columns of the current row into the fields of the Task struct
	err = row.Scan(&task.Id, &task.CreatedAt, &task.UpdatedAt, &task.Item, &task.Done, &task.DoneAt)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			// Handle case where no rows were returned
			return &entity.Task{}, fmt.Errorf("no taskusecase found with ID %s", task.Id)
		}
		// Handle other errors
		return &entity.Task{}, err
	}

	return &task, nil

}

func (r *TaskRepository) Remove(eTask *entity.Task) (string, error) {

	tx, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(r.qry.Delete)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	result, err := stmt.Exec(eTask.Id)
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

	return eTask.Id, nil
}
