package handlers

import (
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

	return c.String(http.StatusOK, requestBody.Nickname)
}
