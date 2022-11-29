package users

import (
	"fmt"

	"github.com/BFDavidGamboa/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewRequestedRangeNotSatisfiable(fmt.Sprintf("user %d not found", user.Id))
	}

	*user = *result

	return nil
}

func (user *User) Save() *errors.RestErr {

	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exist", user.Id))
	}
	usersDB[user.Id] = user
	return nil
}
