package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", welcome)
	e.GET("/user/sign-in", handleSignIn)
	e.GET("/user/sign-up", handleSignUp)

	e.Logger.Fatal(e.Start(":3000"))
}

func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to myapp")
}

func handleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, "Welcome to signin")
}

func handleSignUp(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to myapp")
}
