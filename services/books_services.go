package services

import (
	"github.com/BFDavidGamboa/bookstore_users-api/domain/books"
	"github.com/BFDavidGamboa/bookstore_users-api/utils/errors"
)

func GetBook(bookId int64) (*books.Book, *errors.RestErr) {
	if bookId <= 0 {
		return nil, errors.NewBadRequestError("invalid book Id")
	}

	result := &books.Book{Id: bookId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateBook(book books.Book) (*books.Book, *errors.RestErr) {
	if err := book.Validate(); err != nil {
		return nil, err
	}

	if err := book.Save(); err != nil {
		return nil, err
	}
	return &book, nil
}
