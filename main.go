package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/renxzen/geth/cmd/database"
	home "github.com/renxzen/geth/cmd/home/api"
	todo "github.com/renxzen/geth/cmd/todo/api"
)

func main() {
	defer db.Client.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/static", "cmd/static")

	home.RegisterUi(e)
	todo.RegisterUi(e)
	todo.RegisterWeb(e)

	e.Logger.Fatal(e.Start(":8000"))
}
