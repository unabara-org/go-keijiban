package handlers

import (
	"net/http"

	"github.com/unabara-org/go-keijiban/data"

	"github.com/labstack/echo"
)

func GetComments(c echo.Context) error {
	commentsRepository := data.NewCommentsRepository()
	comments, err := commentsRepository.Read()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, comments)
}
