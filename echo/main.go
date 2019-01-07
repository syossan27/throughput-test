package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
}

var db *gorm.DB

func handler(c echo.Context) error {
	user := GetUser()
	return c.JSON(http.StatusOK, user)
}

func main() {
	var err error
	db, err = gorm.Open("mysql", "root:@/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(100)

	e := echo.New()
	e.GET("/", handler)
	e.Logger.Fatal(
		e.Start(fmt.Sprintf("%s:%s", "127.0.0.1", "8080")),
	)
}

func GetUser() User {
	user := User{}
	if err := db.First(&user, 1).Error; err != nil {
		log.Fatal(err)
	}
	return user
}

