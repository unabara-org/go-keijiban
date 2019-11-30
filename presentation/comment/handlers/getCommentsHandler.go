package handlers

import (
	domainComment "github.com/unabara-org/go-keijiban/domain/comment"
	"net/http"

	"github.com/unabara-org/go-keijiban/data"

	"github.com/labstack/echo"
)

type responseComment struct {
	Id string
	Nickname string
	Body string
}

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

	var responseComments []responseComment

	for _, comment := range comments {
		responseComments = append(responseComments, mapCommentToResponseComment(comment))
	}

	return c.JSON(http.StatusOK, responseComments)
}

func mapCommentToResponseComment(comment domainComment.Comment) responseComment {
	return responseComment{
		Id:       comment.Id.String(),
		Nickname: comment.Nickname,
		Body:     comment.Body,
	}
}
