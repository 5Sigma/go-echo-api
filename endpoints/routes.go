package endpoints

import "github.com/labstack/echo"

// Setup - Sets up endpoint routing
func Setup(e *echo.Echo, h Handler) {
	secured := e.Group("")
	h.SecureGroup(secured)

	secured.GET("/users", h.ListUsers)
	secured.GET("/users/:id", h.GetUser)
	secured.GET("/widgets", h.ListWidgets)
}
