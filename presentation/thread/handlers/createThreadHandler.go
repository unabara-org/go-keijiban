package handlers

import (
	dataNewThread "github.com/unabara-org/go-keijiban/data/new_thread"
	domainComment "github.com/unabara-org/go-keijiban/domain/comment"
	domainNewThread "github.com/unabara-org/go-keijiban/domain/new_thread"
	domainThread "github.com/unabara-org/go-keijiban/domain/thread"
	"net/http"

	"github.com/labstack/echo"
	"github.com/unabara-org/go-keijiban/data"
)

type createRequestBody struct {
	Title    string `json:"title"`
	Nickname string `json:"nickname"`
	Body     string `json:"body"`
}

func CreateThread(c echo.Context) error {
	requestBody := new(createRequestBody)

	if err := c.Bind(requestBody); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	comment := domainComment.NewComment(
		domainComment.NewId(),
		requestBody.Nickname,
		requestBody.Body,
	)

	newThread := domainNewThread.NewNewThread(
		domainThread.NewId(),
		requestBody.Title,
		comment,
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

	newThreadRepository := dataNewThread.NewNewThreadRepository(db)

	if err := newThreadRepository.Create(newThread); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusCreated)
}
