package api

import (
	"context"

	"github.com/labstack/echo/v4"

	"github.com/renxzen/geth/cmd/home/views"
)

func RegisterUi(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		component := views.Home()
		return component.Render(context.Background(), c.Response().Writer)
	})
}
