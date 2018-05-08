package endpoints

import "github.com/labstack/echo"

// Setup - Sets up endpoint routing
func Setup(e *echo.Echo, h Handler) {
	api := e.Group("/api")

	api.GET("/users", h.ListUsers)
	api.GET("/user/:id", h.GetUser)
}
