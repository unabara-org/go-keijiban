package handlers

import (
	"github.com/labstack/echo"
	"github.com/unabara-org/go-keijiban/data"
	domainComment "github.com/unabara-org/go-keijiban/domain/comment"
	"github.com/unabara-org/go-keijiban/presentation/comment/mappers"
	"net/http"
)

type updateRequestBody struct {
	Body     string `json:"body"`
}

func UpdateComment(c echo.Context) error {
	db, err := data.NewDatabase()

	if err != nil {
		return err
	}

	defer func() {
		if err := db.Close(); err != nil {
			// error handle
		}
	}()

	id, err := domainComment.NewIdFromString(c.Param("id"))

	if err != nil {
		return err
	}

	commentsRepository := data.NewCommentsRepository(db)

	comment, err := commentsRepository.Find(*id)

	if err != nil {
		return err
	}

	requestBody := new(updateRequestBody)

	if err := c.Bind(requestBody); err != nil {
		return err
	}

	comment.UpdateBody(requestBody.Body)

	if err := commentsRepository.Update(*comment); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, mappers.MapCommentToResponseComment(*comment))
}

