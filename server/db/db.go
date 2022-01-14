package db

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/martel-darin/graphql_app/graph/model"
)

var (
	ConnectionPool *pgx.Conn
)

func InitDB() error {
	var err error
	ConnectionPool, err = pgx.Connect(context.Background(), "postgres://postgres:password@localhost:5432/postgres")
	if err != nil {
		return err
	}
	return nil
}

func InsertUser(ctx context.Context, user *model.User) error {
	command := `INSERT INTO users(id, name)
	VALUES($1, $2)`

	if _, err := ConnectionPool.Exec(
		ctx,
		command,
		user.ID,
		user.Name,
	); err != nil {
		return err
	}
	return nil
}

func InsertTodo(ctx context.Context, todo *model.Todo) error {
	insertCommand := `INSERT INTO todos(id, text, done, user_id)
	VALUES($1, $2, $3, $4)`

	queryCommand := `SELECT id, name FROM users WHERE id = $1`

	tx, err := ConnectionPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(
		ctx,
		insertCommand,
		todo.ID,
		todo.Text,
		todo.Done,
		todo.User.ID,
	); err != nil {
		return err
	}

	var user model.User
	row := tx.QueryRow(ctx, queryCommand, todo.User.ID)
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		return err
	} else {
		todo.User = &user

	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func FetchUsers(ctx context.Context) ([]*model.User, error) {

	command := `SELECT id, name FROM users`
	rows, err := ConnectionPool.Query(
		ctx,
		command,
	)
	if err != nil {
		return nil, err
	}

	users := make([]*model.User, 0)
	for rows.Next() {
		tmp := new(model.User)
		rows.Scan(&tmp.ID, &tmp.Name)
		users = append(users, tmp)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return users, nil
}

func FetchTodos(ctx context.Context) ([]*model.Todo, error) {
	command := `SELECT todos.id, todos.text, todos.done, todos.user_id, users.name 
	FROM todos JOIN users ON todos.user_id = users.id`
	rows, err := ConnectionPool.Query(
		ctx,
		command,
	)
	if err != nil {
		return nil, err
	}

	todos := make([]*model.Todo, 0)
	for rows.Next() {
		tmp := new(model.Todo)
		tmp.User = new(model.User)
		rows.Scan(&tmp.ID, &tmp.Text, &tmp.Done, &tmp.User.ID, &tmp.User.Name)
		todos = append(todos, tmp)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return todos, nil
}
