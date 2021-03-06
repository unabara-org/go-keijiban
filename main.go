package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	commentHandlers "github.com/unabara-org/go-keijiban/presentation/comment/handlers"
	threadHandlers "github.com/unabara-org/go-keijiban/presentation/thread/handlers"
	"log"
	"net/http"
)

func main() {
	envLoad()

	e := echo.New()

	//logger, dispose := infrastructure.NewFileLogger()
	//defer (func () {
	//	if err := dispose(); err != nil {
	//		// なんかエラーハンドリング
	//	}
	//})()

	//e.Logger = logger

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/threads", threadHandlers.CreateThread)

	e.GET("/comments", commentHandlers.GetComments)

	e.POST("/comments", commentHandlers.CreateComment)

	e.PATCH("/comments/:id", commentHandlers.UpdateComment)

	e.DELETE("/comments/:id", commentHandlers.DeleteComment)

	_ = e.Start(":1323")
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
