package main

import (
	"github.com/fukuyama012/go-server-layered-sample/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	router.Init(e)
	e.Logger.Fatal(e.Start(":1323"))
}
