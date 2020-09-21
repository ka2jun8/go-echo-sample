package main

import (
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	router := Router()

	// Start server
	router.Logger.Fatal(router.Start(":1323"))
}

// Router is ...
func Router() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/healthz", healthz)

	e.GET("/count", getCount)
	e.POST("/count", postCount)

	return e
}

// health check Handler
func healthz(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

var count int

// get count Handler
func getCount(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprint(count))
}

// post count Handler
func postCount(c echo.Context) error {
	count++
	return c.NoContent(http.StatusOK)
}
