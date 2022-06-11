package users

import (
	"fmt"
	"strings"

	"github.com/viizgar/go-bookstore_users-api/datasources/db/users_db"
	"github.com/viizgar/go-bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = `INSERT INTO users(first_name, last_name, email, date_created) VALUES ($1,$2,$3,$4) RETURNING id;`
	queryGetUser    = `SELECT id, first_name, last_name, email, date_created FROM users WHERE id=$1;`
)

func (user *User) Get() *errors.Error {

	err := users_db.Client.QueryRow(queryGetUser, user.Id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}

		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get an user: %s", err.Error()))
	}

	return nil
}

func (user *User) Save() *errors.Error {

	var id int64
	err := users_db.Client.QueryRow(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated).Scan(&id)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save an user: %s", err.Error()))
	}

	user.Id = id
	return nil
}
