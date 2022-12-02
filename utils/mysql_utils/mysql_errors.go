package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/BFDavidGamboa/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	fmt.Println(sqlErr.Message)
	switch sqlErr.Number {
	//1292
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}
