package books

import (
	"fmt"

	"github.com/BFDavidGamboa/bookstore_users-api/utils/date_utils"
	"github.com/BFDavidGamboa/bookstore_users-api/utils/errors"
)

var (
	booksDB = make(map[int64]*Book)
)

func (book *Book) Get() *errors.RestErr {
	result := booksDB[book.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("book %v not found", book.Id))
	}

	*book = *result
	return nil
}

func (book *Book) Save() *errors.RestErr {

	current := booksDB[book.Id]
	if current != nil {
		if current.Tittle == book.Tittle {
			return errors.NewBadRequestError(fmt.Sprintf("book %s already exists", book.Tittle))
		}
	}

	book.DateCreated = date_utils.GetNowString()

	booksDB[book.Id] = book
	return nil
}
