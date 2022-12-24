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
	Update(todo *models.Todo) (model models.Todo, err error)
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

// Update handle an action to update the todo data
func (t *todoRepository) Update(todo *models.Todo) (model models.Todo, err error) {
	var sql = `
		SELECT
			id, user_id, todo, date, is_finished, created_at, updated_at
		FROM
			todos
		WHERE user_id = $1 AND id = $2`

	err = t.db.QueryRow(context.Background(), sql,
		todo.UserId, todo.Id,
	).Scan(
		&model.Id, &model.UserId, &model.Todo, &model.Date,
		&model.IsFinished, &model.CreatedAt, &model.UpdatedAt,
	)
	if err != nil {
		return
	}

	var updateSql = `
		UPDATE
			todos
		SET
			todo = $1, date = $2, is_finished = $3,
			updated_at = $4
		WHERE
			id = $5 AND user_id = $6`

	// make updated data
	model.Todo = todo.Todo
	model.Date = todo.Date
	model.IsFinished = todo.IsFinished
	model.UpdatedAt = time.Now().Unix()

	_, err = t.db.Exec(context.Background(), updateSql,
		model.Todo, model.Date, model.IsFinished, model.UpdatedAt, model.Id, model.UserId,
	)
	return
}
