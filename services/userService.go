package services

import (
	"github.com/viizgar/go-bookstore_users-api/domain/users"
	"github.com/viizgar/go-bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.Error) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.Error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(user users.User) (*users.User, *errors.Error) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email

	if err := current.Validate(); err != nil {
		return nil, err
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}
