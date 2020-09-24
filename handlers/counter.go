package handlers

import (
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

var count int

// get count Handler
func GetCountHandler(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprint(count))
}

// post count Handler
func PostCountHandler(c echo.Context) error {
	count++
	return c.NoContent(http.StatusOK)
}
