package users

import (
	"fmt"
	"strings"

	"github.com/BFDavidGamboa/bookstore_users-api/datasources/mysql/users_db"
	"github.com/BFDavidGamboa/bookstore_users-api/utils/date_utils"
	"github.com/BFDavidGamboa/bookstore_users-api/utils/errors"
)

const (
	errorNoRows      = "no rows in result set"
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, password, date_created) VALUES (?, ?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, password, date_created FROM users WHERE id = ?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get user %d : %s", user.Id, err.Error()),
		)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.DateCreated)

	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(
				fmt.Sprintf("email %s already exists", user.Email),
			)
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error trying to create user %s", err.Error()),
		)
	}

	// another aproach could be using Client.Exec which not prepare
	// the query and execute it instantly
	// result, err := users_db.Client.Exec(querInsertUser, user.FirstName, user.LastName, user.Email, user.Password, user.DateCreated)

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()),
		)
	}

	user.Id = userId
	return nil
}
