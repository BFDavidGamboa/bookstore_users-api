package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	Client *sql.DB
)

func init() {
	var err error

	err = godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//"%s:%s@tcp(%s)/%s?charser=utf8" represents
	//%s user name: %s password @ tcp ( %s host)/ %s sql scheme ? charset
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		os.Getenv("DB_MYSQL_USERNAME"),
		os.Getenv("DB_MYSQL_PASSWORD"),
		os.Getenv("DB_MYSQL_HOST"),
		os.Getenv("DB_MYSQL_SCHEME"),
	)

	Client, err = sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
