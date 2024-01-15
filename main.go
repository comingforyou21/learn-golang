package main

import (
	"context"
	"fmt"
	"mymodule/db"
	"mymodule/handler"
	log "mymodule/log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func init() {
	fmt.Println("init package main")
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

// test
func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "changeme",
		DbName:   "demo",
	}
	sql.Connect()
	defer sql.Close()

	var email string
	err := sql.Db.GetContext(context.Background(), &email, "SELECT email FROM users WHERE email=$1", "abc@gmail.com")
	if err != nil {
		log.Error(err.Error())
	}

	e := echo.New()
	e.GET("/", welcome)
	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)

	e.Logger.Fatal(e.Start(":3000"))
}

func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to myapp")
}
