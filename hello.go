package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Hello(c echo.Context) error {
	// Write
	go func() {
		C1 <- "Hello が実行されました！"
	}()
	return c.String(http.StatusOK, "Hello, World!")
}
