package users

import (
	"errors"
	"fmt"
	"strings"

	"github.com/BFDavidGamboa/bookstore_users-api/datasources/mysql/users_db"
	"github.com/BFDavidGamboa/bookstore_users-api/utils/mysql_utils"
	"github.com/BFDavidGamboa/bookstore_utils-go/logger"
	"github.com/BFDavidGamboa/bookstore_utils-go/rest_errors"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, password, status) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=?, status=?, password=? WHERE id = ?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindByStatus     = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?;"
	queryEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email = ? AND password=? AND status=?;"
)

func (user *User) Get() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {

		logger.Error("error when trying to prepare get user statement", err)
		return rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when trying to prepare get user statement", getErr)
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare user statement", err)
		return rest_errors.NewInternalServerError("error when trying to save user", errors.New("database error"))
	}

	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)
	if saveErr != nil {
		logger.Error("error when trying to user", saveErr)
		return rest_errors.NewInternalServerError("error when trying to save user", errors.New("database error"))
	}

	// another aproach could be using Client.Exec which not prepare
	// the query and execute it instantly
	// result, err := users_db.Client.Exec(querInsertUser, user.FirstName, user.LastName, user.Email, user.Password, user.DateCreated)

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error(" error when trying to get last insert id after creating a new user", saveErr)
		return rest_errors.NewInternalServerError("error when trying to save user", errors.New("database error"))
	}

	user.Id = userId
	return nil
}

func (user *User) Update() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare query to update user", err)
		return rest_errors.NewInternalServerError("error when trying to update user", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.FirstName, user.LastName, user.Email, user.Status, user.Password)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare query to delete user", err)
		return rest_errors.NewInternalServerError("error when trying to delete user", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		logger.Error("error when trying to delete user", err)
		return rest_errors.NewInternalServerError("error when trying to delete user", errors.New("database error"))
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, rest_errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error when trying to prepare query to find users", err)
		return nil, rest_errors.NewInternalServerError("error when trying to find user", errors.New("database errors"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find users", err)
		return nil, rest_errors.NewInternalServerError("error when trying to find user", errors.New("database errors"))
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when scan user row int user users", err)
			return nil, rest_errors.NewInternalServerError("error when trying to find user", errors.New("database errors"))
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}

func (user *User) FindByEmailAndPassword() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get emayl and password statement", err)
		return rest_errors.NewInternalServerError("error when trying to find by email user", errors.New("database error"))
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Email, user.Password, StatusActive)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return rest_errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to prepare get user by email and password", getErr)
		return mysql_utils.ParseError(getErr)
	}
	return nil
}
