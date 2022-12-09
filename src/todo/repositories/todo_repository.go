package repositories

import (
	"context"
	"go_todo_api/src/shared/databases"
	"go_todo_api/src/todo/models"
	"time"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// TodoRepository is an abstract that contains
// methods to manipulate todo data
//
//counterfeiter:generate . TodoRepository
type TodoRepository interface {
	Store(todo *models.Todo) error
	Find(userId int64) (todos []models.Todo, err error)
	Detail(userId, todoId int64) (todo models.Todo, err error)
}

// todoRepository is a struct that has methods
// to manipulate todo data
type todoRepository struct {
	db databases.DB
}

// NewTodoRepository creates an instance of todoRepository
func NewTodoRepository(db databases.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

// Store handle an action to store todo data
func (t *todoRepository) Store(todo *models.Todo) error {
	var sql = `
		INSERT INTO todos
		(
			user_id, todo, date, is_finished, created_at, updated_at
		)
		VALUES
		(
			$1, $2, $3, $4, $5, $6
		) RETURNING id, user_id, todo, date, is_finished, created_at, updated_at`

	// create timestamps
	timestamps := time.Now().Unix()
	todo.CreatedAt = timestamps
	todo.UpdatedAt = timestamps

	err := t.db.QueryRow(context.Background(), sql,
		&todo.UserId, &todo.Todo, &todo.Date, &todo.IsFinished, &todo.CreatedAt, &todo.UpdatedAt).Scan(
		&todo.Id, &todo.UserId, &todo.Todo, &todo.Date, &todo.IsFinished, &todo.CreatedAt, &todo.UpdatedAt)

	return err
}

// Find handle an action to provide user todo datas
// by the given user id
func (t *todoRepository) Find(userId int64) (todos []models.Todo, err error) {
	var sql = `
		SELECT
			id, todo, date, is_finished, created_at, updated_at
		FROM
			todos
		WHERE
			user_id = $1`

	rows, err := t.db.Query(context.Background(), sql, userId)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&todo.Id, &todo.Todo, &todo.Date, &todo.IsFinished, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return
		}

		todos = append(todos, todo)
	}
	return
}

// Detail handle an action to provide detail of user todo data
func (t *todoRepository) Detail(userId, todoId int64) (todo models.Todo, err error) {
	var sql = `
		SELECT
			id, todo, date, is_finished, created_at, updated_at
		FROM
			todos
		WHERE user_id = $1 AND id = $2`

	err = t.db.QueryRow(context.Background(), sql, userId, todoId).Scan(
		&todo.Id, &todo.Todo, &todo.Date, &todo.IsFinished, &todo.CreatedAt, &todo.UpdatedAt)
	return
}
