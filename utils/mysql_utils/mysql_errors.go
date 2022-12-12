package mysql_utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/BFDavidGamboa/bookstore_utils-go/rest_errors"
	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewNotFoundError("no record matching id")
		}
		return rest_errors.NewInternalServerError("error when trying to get user", errors.New("error parsing database response"))
	}
	fmt.Println(sqlErr.Message)
	switch sqlErr.Number {
	//1292
	case 1062:
		return rest_errors.NewBadRequestError("invalid data")
	}
	return rest_errors.NewInternalServerError("error when trying to get user", errors.New("error processing request"))
}
