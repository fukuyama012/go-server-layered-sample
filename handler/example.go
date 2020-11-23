package handler

import (
	"net/http"
	"strconv"

	"github.com/fukuyama012/go-server-layered-sample/usecase"
	"github.com/labstack/echo/v4"
)

type (
	// Example interface
	Example interface {
		Get(c echo.Context) error
	}

	// ExampleImpl struct
	ExampleImpl struct {
		usecase usecase.Example
	}
)

// NewExampleHandler returns example handler instance
func NewExampleHandler() Example {
	return &ExampleImpl{
		usecase: initExample("hint"),
	}
}

// Get get example
func (h *ExampleImpl) Get(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	res, err := h.usecase.Get(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"example": res})
}
