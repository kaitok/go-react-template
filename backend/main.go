package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "jon@lafstack.com",
		}
		return c.JSON(http.StatusOK, u)
	})

	e.Start(":8080")
}
