package endpoints

import "github.com/labstack/echo"

// Setup - Sets up endpoint routing
func Setup(e *echo.Echo, h Handler) {
	api := e.Group("/api")
	h.SecureGroup(api)

	api.GET("/users", h.ListUsers)
	api.GET("/users/:id", h.GetUser)
}
