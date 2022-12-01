package books

import (
	"fmt"
	"strings"

	"github.com/BFDavidGamboa/bookstore_users-api/datasources/mysql/users_db"
	"github.com/BFDavidGamboa/bookstore_users-api/utils/date_utils"
	"github.com/BFDavidGamboa/bookstore_users-api/utils/errors"
)

const (
	errorNoRows     = "no rows in result set"
	indexUniqueIsbn = "isbn_UNIQUE"
	queryInsertBook = "INSERT INTO books(tittle, author, country, isbn, date_created) VALUES (?, ?, ?, ?, ?);"
	queryGetBook    = "SELECT id, tittle, author, country, isbn, date_created FROM books where id = ?;"
)

func (book *Book) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetBook)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(book.Id)

	if err := result.Scan(&book.Id, &book.Tittle, &book.Author, &book.Country, &book.Isbn, &book.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(
				fmt.Sprintf("book %d not found", book.Id),
			)
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get book %d: %v", book.Id, err.Error()),
		)
	}

	return nil
}

func (book *Book) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertBook)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	book.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(book.Tittle, book.Author, book.Country, book.Isbn, book.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueIsbn) {
			return errors.NewBadRequestError(
				fmt.Sprintf("book %s already exists", book.Isbn),
			)
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error trying to create book %s", err.Error()),
		)
	}

	bookId, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save book %s", err.Error()),
		)
	}

	book.Id = bookId
	return nil
}
