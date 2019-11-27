package main

import (
	"github.com/joho/godotenv"
	commentHandlers "github.com/unabara-org/go-keijiban/presentation/comment/handlers"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	envLoad()

	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/", commentHandlers.CreateComment)

	_ = e.Start(":1323")
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
