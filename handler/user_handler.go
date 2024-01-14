package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "Anhnh",
		"email": "abc@gmail.com",
	})
}

func HandleSignUp(c echo.Context) error {
	type User struct {
		Email    string `json:"email"`
		FullName string `json:"name"`
		Age      int    `json:"age"`
	}
	user := User{
		Email:    "anhnh11111@gmail.com",
		FullName: "anhnh",
		Age:      30,
	}
	return c.JSON(http.StatusOK, user)
}
