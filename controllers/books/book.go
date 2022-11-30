package books

import (
	"net/http"
	"strconv"

	"github.com/BFDavidGamboa/bookstore_users-api/domain/books"
	"github.com/BFDavidGamboa/bookstore_users-api/services"
	"github.com/BFDavidGamboa/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book books.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	result, saveErr := services.CreateBook(book)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetBook(c *gin.Context) {
	bookId, bookErr := strconv.ParseInt(c.Param("book_id"), 10, 64)
	if bookErr != nil {
		err := errors.NewBadRequestError("book id should be a number")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	book, getErr := services.GetBook(bookId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, book)
}

func FindBook(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "implement me!")
}
