package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/renxzen/geth/cmd/todo/domain"
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
		todos, err := domain.ListAll()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		component := views.List(todos)
		return component.Render(context.Background(), c.Response().Writer)
	})

	r.POST("", func(c echo.Context) error {
		content := c.FormValue("content")
		if content == "" {
			return c.String(http.StatusBadRequest, "Content cannot be empty")
		}

		todo := domain.Todo{
			Content:   content,
			Completed: false,
		}

		todo, err := domain.CreateTodo(todo)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		component := views.Todo(todo)
		return component.Render(context.Background(), c.Response().Writer)
	})

	r.PUT("/:id", func(c echo.Context) error {
		paramId := c.Param("id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid id")
		}

		todo, err := domain.ToggleCompletedById(int64(id))
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		component := views.Todo(todo)
		return component.Render(context.Background(), c.Response().Writer)
	})

	r.DELETE("/:id", func(c echo.Context) error {
		paramId := c.Param("id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid id")
		}

		err = domain.DeleteById(int64(id))
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	})
}
