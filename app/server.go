package main

import (
	"echo_app/app/handler"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, Echo!!")
	})

	// routing
	e.GET("/users", handler.GetAll())
	e.GET("/users/:id", handler.Get())
	e.POST("/users", handler.Create())
	e.PUT("/users/:id", handler.Update())
	e.DELETE("/users/:id", handler.Delete())

	// start
	e.Logger.Fatal(e.Start(":8080"))
}
