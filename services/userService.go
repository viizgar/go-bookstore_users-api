package services

import (
	"github.com/viizgar/go-bookstore_users-api/domain/users"
	"github.com/viizgar/go-bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.Error) {
	return &user, nil
}
