package db

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/martel-darin/graphql_app/models"
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

func InsertUser(user *models.User) error {
	command := `INSERT INTO USERS(USERNAME, PASSWORD, EMAIL, CREATED_ON)
	VALUES($1, $2, $3, $4)`

	if _, err := ConnectionPool.Exec(
		context.Background(),
		command,
		user.Username,
		user.Password,
		user.Email,
		user.CreatedOn,
	); err != nil {
		return err
	}

	return nil
}

func FetchUsers() ([]models.User, error) {
	users := make([]models.User, 0)

	command := `SELECT * FROM users`
	rows, err := ConnectionPool.Query(
		context.Background(),
		command,
	)
	if err != nil {
		return users, err
	}

	var tmp models.User
	for rows.Next() {
		rows.Scan(&tmp.ID, &tmp.Username, &tmp.Password, &tmp.Email, &tmp.CreatedOn)
		users = append(users, tmp)
	}

	if rows.Err() != nil {
		return users, rows.Err()
	}
	return users, nil
}
