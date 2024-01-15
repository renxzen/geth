package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/renxzen/geth/cmd/todo/domain"
)

func RegisterWeb(e *echo.Echo) {
	r := e.Group("/todo")

	r.GET("", func(c echo.Context) error {
		todos, err := domain.ListAll()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, todos)
	})

	r.POST("", func(c echo.Context) error {
		var todo domain.Todo
		if err := c.Bind(&todo); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		todo.Completed = false
		todo, err := domain.CreateTodo(todo)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, todo)
	})

	r.GET("/:id", func(c echo.Context) error {
		paramId := c.Param("id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid id")
		}

		todo, err := domain.GetById(int64(id))
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, todo)
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

		return c.JSON(http.StatusOK, todo)
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

