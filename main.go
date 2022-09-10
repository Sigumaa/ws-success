package main

import (
	_ "github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// C1 is a channel
var C1 = make(chan string)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/ws", ServeWs)
	e.GET("/hello", Hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
