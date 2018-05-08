package endpoints

import (
	"strconv"

	"github.com/5sigma/go-echo-api/models"

	"github.com/labstack/echo"
)

// GetUser - Retrun the information for a user
func (h *Handler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := h.DB.GetUserByID(uint(id))
	return c.JSON(200, user)
}

// ListUsers - Retrun the information for all users
func (h *Handler) ListUsers(c echo.Context) error {
	users := h.DB.AllUsers()
	response := struct {
		Users []models.User `json:"users"`
	}{users}
	return c.JSON(200, response)
}
