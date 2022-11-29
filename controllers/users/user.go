package users

import (
	"fmt"
	"net/http"

	"github.com/BFDavidGamboa/bookstore_users-api/domain/user"
	"github.com/BFDavidGamboa/bookstore_users-api/services"
	"github.com/BFDavidGamboa/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user user.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	result, saveErr := services.CreateUser(&user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "implement me!")
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me!")
// }

func UpdateUser() {}

func DeleteUser() {}
