package main

import (
	_ "github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var C1 = make(chan string)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/ws", Connect)
	e.GET("/hello", Hello)
	e.Logger.Fatal(e.Start(":1323"))
}
