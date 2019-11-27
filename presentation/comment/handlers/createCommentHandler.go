package handlers

import (
	"net/http"

	"github.com/unabara-org/go-keijiban/data"
	"github.com/unabara-org/go-keijiban/domain"

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

	comment := domain.NewComment(
		domain.NewCommentId(),
		requestBody.Nickname,
		requestBody.Body,
	)

	db, err := data.NewDatabase()

	if err != nil {
		return err
	}

	// https://github.com/go-sql-driver/mysql/issues/910 を参考にした
	defer func() {
		if err := db.Close(); err != nil {
			// error handle
		}
	}()

	commentsRepository := data.NewCommentsRepository(db)

	if err := commentsRepository.Create(comment); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, comment)
}
