package api

import (
	"context"

	"github.com/labstack/echo/v4"

	"github.com/renxzen/geth/cmd/todo/views"
)

func RegisterUi(e *echo.Echo) {
	r := e.Group("/todo/ui")

	r.GET("", func(c echo.Context) error {
		component := views.Main()
		return component.Render(context.Background(), c.Response().Writer)
	})

	r.GET("/header", func(c echo.Context) error {
		component := views.Header()
		return component.Render(context.Background(), c.Response().Writer)
	})

	r.GET("/list", func(c echo.Context) error {
		component := views.List()
		return component.Render(context.Background(), c.Response().Writer)
	})
}
