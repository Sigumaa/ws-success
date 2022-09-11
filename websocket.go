package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ChangeVolume(c echo.Context) error {
	go func() {
		C1 <- "ChangeVolume が実行されました！"
	}()
	return c.String(http.StatusOK, "ChangeVolume, World!")
}
