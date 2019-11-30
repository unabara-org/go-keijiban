package handlers

import (
	"net/http"

	domainComment "github.com/unabara-org/go-keijiban/domain/comment"

	"github.com/labstack/echo"
	"github.com/unabara-org/go-keijiban/data"
)

func DeleteComment(c echo.Context) error {
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

	id, err := domainComment.NewIdFromString(c.Param("id"))

	if err != nil {
		return err
	}

	commentsRepository := data.NewCommentsRepository(db)
	comment, err := commentsRepository.Find(*id)

	if err != nil {
		return err
	}

	if err := commentsRepository.Delete(*comment); err != nil {
		return err
	}

	return c.JSON(http.StatusNoContent, nil)
}
