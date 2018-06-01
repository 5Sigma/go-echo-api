package endpoints

import (
	"github.com/5sigma/go-echo-api/dao"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Handler - API Request handler
type Handler struct {
	DB *dao.DAO
}

func (h *Handler) secureGroupValidator(key string, c echo.Context) (bool, error) {
	apiKey := h.DB.LookupAPIKey(key)
	if apiKey == nil {
		return false, nil
	}
	c.Set("APIKey", apiKey)
	currentUser := h.DB.GetUserByID(apiKey.UserID)
	c.Set("CurrentUser", currentUser)
	return true, nil
}

// SecureGroup - Secure a group using an API key.
func (h *Handler) SecureGroup(g *echo.Group) {
	g.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup:  "header:Authorization",
		AuthScheme: "Bearer",
		Validator:  h.secureGroupValidator,
	}))
}
