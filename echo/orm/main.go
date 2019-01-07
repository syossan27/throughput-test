package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

var db *sql.DB

func handler(c echo.Context) error {
	user := GetUser()
	return c.JSON(http.StatusOK, user)
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)

	e := echo.New()
	e.GET("/", handler)
	e.Logger.Fatal(
		e.Start(fmt.Sprintf("%s:%s", "127.0.0.1", "8080")),
	)
}

func GetUser() User {
	user := User{}
	if err := db.QueryRow("SELECT name FROM users WHERE id = ?", 1).Scan(&user.Name); err != nil {
		log.Fatal(err)
	}

	return user
}

