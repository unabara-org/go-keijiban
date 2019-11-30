package handlers

import (
	"net/http"

	presentationComment "github.com/unabara-org/go-keijiban/presentation/comment"
	"github.com/unabara-org/go-keijiban/presentation/comment/mappers"

	"github.com/unabara-org/go-keijiban/data"

	"github.com/labstack/echo"
)

func GetComments(c echo.Context) error {
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
	comments, err := commentsRepository.Read()

	if err != nil {
		return err
	}

	var responseComments []presentationComment.ResponseComment

	for _, comment := range comments {
		responseComments = append(responseComments, mappers.MapCommentToResponseComment(comment))
	}

	return c.JSON(http.StatusOK, responseComments)
}
