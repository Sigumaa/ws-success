package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/ws", ServeWs)

	// Send message to all clients
	// e.GET("/hello", Hello)

	// Send JSON to all clients
	e.GET("/changeVolume", ChangeVolume)
	e.GET("/memoUpdate", MemoUpdate)
	e.GET("/speakerChange", SpeakerChange)
	e.GET("/sceneChange", SceneChange)
	e.GET("/messageUpdate", MessageUpdate)
	e.GET("/Comments", Comments)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
