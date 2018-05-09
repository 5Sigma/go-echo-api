package endpoints

import (
	"errors"

	"github.com/5sigma/go-echo-api/dao"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Handler - API Request handler
type Handler struct {
	DB *dao.DAO
}

// SecureGroup - Secure a group using an API key.
func (h *Handler) SecureGroup(g *echo.Group) {
	g.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup:  "header:Authorization",
		AuthScheme: "Bearer",
		Validator: func(key string, c echo.Context) (bool, error) {
			apiKey := h.DB.LookupAPIKey(key)
			if apiKey == nil {
				return false, errors.New("Invalid API key")
			}
			c.Set("APIKey", apiKey)
			return true, nil
		},
	}))
}
