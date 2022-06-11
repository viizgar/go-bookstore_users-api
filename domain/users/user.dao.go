package users

import (
	"fmt"

	"github.com/viizgar/go-bookstore_users-api/datasources/db/users_db"
	"github.com/viizgar/go-bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = `INSERT INTO users(first_name, last_name, email, date_created) VALUES ($1,$2,$3,$4) RETURNING id;`
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.Error {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

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
