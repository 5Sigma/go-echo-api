package endpoints

import (
	"errors"
	"strconv"

	"github.com/5sigma/go-echo-api/models"
	"github.com/labstack/echo"
)

// GetWidget - Endpoint for GET/widgets/:id
func (h *Handler) GetWidget(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	widget := h.DB.GetWidgetByID(uint(id))
	if widget == nil {
		return errors.New("Widget not found")
	}
	return c.JSON(200, widget)
}

// ListWidgets - Endpoint for GET /widgets
func (h *Handler) ListWidgets(c echo.Context) error {
	user := c.Get("CurrentUser").(*models.User)
	widgets := h.DB.AllWidgetsForUser(user)
	response := struct {
		Widgets []models.Widget `json:"widgets"`
	}{
		Widgets: widgets,
	}
	return c.JSON(200, response)
}

// CreateWidget - Endpoint for POST /widgets
func (h *Handler) CreateWidget(c echo.Context) error {
	widget := new(models.Widget)
	if err := c.Bind(widget); err != nil {
		return err
	}
	user := c.Get("user").(*models.User)
	widget.CreatorID = user.ID
	widget = h.DB.CreateWidget(widget)
	widget.Creator = user
	return c.JSON(200, widget)
}

// DeleteWidget - Endpoint for DELETE /widgets/:id
func (h *Handler) DeleteWidget(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	h.DB.DeleteWidgetByID(uint(id))
	return c.String(200, "")
}

// UpdateWidget - Endpoint for PUT /widgets/:id
func (h *Handler) UpdateWidget(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	widget := new(models.Widget)
	widget.ID = uint(id)
	if err := c.Bind(widget); err != nil {
		return err
	}
	h.DB.UpdateWidget(widget)
	return c.JSON(200, widget)
}
