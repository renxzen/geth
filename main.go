package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	home "github.com/renxzen/geth/cmd/home/api"
	todo "github.com/renxzen/geth/cmd/todo/api"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/static", "cmd/static")

	home.RegisterUi(e)
	todo.RegisterUi(e)

	e.Logger.Fatal(e.Start(":8000"))
}
