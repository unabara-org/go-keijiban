package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetComments(c echo.Context) error {
	return c.String(http.StatusOK, "foo")
}
