package comments

import (
	"net/http"

	"github.com/labstack/echo"
)

func HandleGetComments(c echo.Context) error {
	return c.String(http.StatusOK, "foo")
}
