package main

import (
	"api/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/healthcheck", handler.Hello())
	e.POST("/login", handler.Login())

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.POST("", handler.Restricted())

	e.Start(":3000")
}
