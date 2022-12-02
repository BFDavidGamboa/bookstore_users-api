package users

import (
	"fmt"

	"github.com/BFDavidGamboa/bookstore_users-api/datasources/mysql/users_db"
	"github.com/BFDavidGamboa/bookstore_users-api/utils/errors"
	"github.com/BFDavidGamboa/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser     = "INSERT INTO users(first_name, last_name, email, date_created, password, status) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser        = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryUpdateUser     = "UPDATE users SET first_name=?, last_name=?, email=?, status=?, password=? WHERE id = ?;"
	queryDeleteUser     = "DELETE FROM users WHERE id=?;"
	queryFindUserStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, saveErr := users_db.Client.Prepare(queryInsertUser)
	if saveErr != nil {
		return errors.NewInternalServerError(saveErr.Error())
	}

	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	// another aproach could be using Client.Exec which not prepare
	// the query and execute it instantly
	// result, err := users_db.Client.Exec(querInsertUser, user.FirstName, user.LastName, user.Email, user.Password, user.DateCreated)

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.FirstName, user.LastName, user.Email, user.Status, user.Password)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
