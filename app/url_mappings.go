package app

import (
	"github.com/BFDavidGamboa/bookstore_users-api/controllers/books"
	"github.com/BFDavidGamboa/bookstore_users-api/controllers/ping"
	"github.com/BFDavidGamboa/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("internal/users/search", users.Search)
	router.POST("users/login", users.Login)

	router.POST("/books", books.CreateBook)
	router.GET("/books/:book_id", books.GetBook)
	router.GET("/books/search", books.FindBook)
}
