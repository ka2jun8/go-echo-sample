package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// HealthCheckHandler is health check
func HealthCheckHandler(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
