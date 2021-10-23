package main

import (
	"github.com/zalfrie/go-reminder/handlers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/health-check", handlers.HealthCheck)
	//Authentication routes
	g := e.Group("/v1")
	g.POST("/login", handlers.Login)
	g.POST("/logout", handlers.Logout)

	e.Logger.Fatal(e.Start(":8080"))
}