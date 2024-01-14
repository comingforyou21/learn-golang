package main

import (
	"mymodule/db"
	"mymodule/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

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

	e := echo.New()
	e.GET("/", welcome)
	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)

	e.Logger.Fatal(e.Start(":3000"))
}

func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to myapp")
}
