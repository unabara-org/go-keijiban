package handlers

import (
	domainComment "github.com/unabara-org/go-keijiban/domain/comment"
	"github.com/unabara-org/go-keijiban/presentation/comment/mappers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/unabara-org/go-keijiban/data"
)

type createRequestBody struct {
	Nickname string `json:"nickname"`
	Body     string `json:"body"`
}

func CreateComment(c echo.Context) error {
	requestBody := new(createRequestBody)

	if err := c.Bind(requestBody); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	comment := domainComment.NewComment(
		domainComment.NewId(),
		requestBody.Nickname,
		requestBody.Body,
	)

	db, err := data.NewDatabase()

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// https://github.com/go-sql-driver/mysql/issues/910 を参考にした
	defer func() {
		if err := db.Close(); err != nil {
			// error handle
		}
	}()

	commentsRepository := data.NewCommentsRepository(db)

	if err := commentsRepository.Create(comment); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, mappers.MapCommentToResponseComment(comment))
}
