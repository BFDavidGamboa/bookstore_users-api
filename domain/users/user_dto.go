package users

import (
	"strings"

	"github.com/BFDavidGamboa/bookstore_users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id" binding:"required"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	DateCreated string `json:"date_created"`
}

func (user User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower((user.Email)))
	if user.Email == "" {
		return errors.NewBadRequestError("email is required")
	}
	return nil
}
