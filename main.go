package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/renxzen/geth/cmd/home/api"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/static", "cmd/static")

	api.RegisterUi(e)

	e.Logger.Fatal(e.Start(":8000"))
}
