package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/renxzen/geth/cmd/todo/domain"
	"github.com/renxzen/geth/cmd/todo/views"
)

var global domain.GlobalState

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
		component := views.List(global.Todos)
		return component.Render(context.Background(), c.Response().Writer)
	})

	r.POST("", func(c echo.Context) error {
		content := c.FormValue("content")
		if content == "" {
			return c.String(http.StatusBadRequest, "Content cannot be empty")
		}

		todo := domain.Todo{
			Id:        int64(len(global.Todos) + 1),
			Content:   content,
			Completed: false,
		}

		global.Todos = append(global.Todos, todo)

		component := views.Todo(todo)
		return component.Render(context.Background(), c.Response().Writer)
	})

	r.PUT("/:id", func(c echo.Context) error {
		paramId := c.Param("id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid id")
		}

		for idx := range global.Todos {
			if global.Todos[idx].Id == int64(id) {
				global.Todos[idx].Completed = !global.Todos[idx].Completed
				component := views.Todo(global.Todos[idx])
				return component.Render(context.Background(), c.Response().Writer)
			}
		}

		return c.String(http.StatusBadRequest, "Todo not found")
	})

	r.DELETE("/:id", func(c echo.Context) error {
		paramId := c.Param("id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid id")
		}

		for idx := range global.Todos {
			if global.Todos[idx].Id == int64(id) {
				global.Todos = append(global.Todos[:idx], global.Todos[idx+1:]...)
				return c.NoContent(http.StatusNoContent)
			}
		}

		return c.String(http.StatusBadRequest, "Todo not found")
	})
}
