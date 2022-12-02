package books

import "github.com/BFDavidGamboa/bookstore_users-api/utils/errors"

type Book struct {
	Id          int64  `json:"id" binding:"required"`
	Tittle      string `json:"tittle"`
	Author      string `json:"author"`
	Country     string `json:"country"`
	Isbn        string `json:"isbn"`
	DateCreated string `json:"date_created"`
}

func (book Book) Validate() *errors.RestErr {
	if book.Tittle == "" {
		return errors.NewBadRequestError("tittle is required")
	}
	return nil
}
