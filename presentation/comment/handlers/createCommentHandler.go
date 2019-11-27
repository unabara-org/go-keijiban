package handlers

import (
	"github.com/unabara-org/go-keijiban/data"
	"github.com/unabara-org/go-keijiban/domain"
	"net/http"

	"github.com/labstack/echo"
)

type CommentRequestBody struct {
	Nickname string `json:"nickname"`
	Body     string `json:"body"`
}

func CreateComment(c echo.Context) error {
	requestBody := new(CommentRequestBody)

	if err := c.Bind(requestBody); err != nil {
		return err
	}

	comment := domain.NewComment(requestBody.Nickname, requestBody.Body)

	commentsRepository := data.NewCommentsRepository()

	if err := commentsRepository.Create(comment); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, comment)
}
