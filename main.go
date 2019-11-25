package main

import (
	commentHandlers "github.com/unabara-org/go-keijiban/presentation/comment/handlers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/", commentHandlers.CreateComment)

	_ = e.Start(":1323")
}
