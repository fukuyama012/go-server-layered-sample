package router

import (
	"github.com/fukuyama012/go-server-layered-sample/handler"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Init init router
func Init(e *echo.Echo) {
	e.GET("/healthcheck", func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json")
		return c.String(http.StatusOK, "OK")
	})

	v1 := e.Group("v1")
	example := handler.NewExampleHandler()
	v1.GET("/example/:id", example.Get)
}
