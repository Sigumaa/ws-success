package main

import (
	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {

	// Write
	// C1 <- "Hello が実行されました！"
	//return c.String(http.StatusOK, "Hello, World!")

	// 実際に使用する際は消す
	return nil
}
