package handlers

import (
	"net/http"

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

	return c.JSON(http.StatusOK, comments)
}
