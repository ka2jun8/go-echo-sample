package server

import (
	"github.com/ka2jun8/go-echo-sample/handlers"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Router is server routes
func Router() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/healthz", handlers.HealthCheckHandler)

	e.GET("/count", handlers.GetCountHandler)
	e.POST("/count", handlers.PostCountHandler)

	e.GET("/files", handlers.InputFormHandler)
	e.POST("/files", handlers.UploadFilesHandler)

	return e
}
