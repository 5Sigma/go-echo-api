package dao

import "github.com/5sigma/go-echo-api/models"

// AllWidgets - Returns all widgets
func (dao *DAO) AllWidgets() []models.Widget {
	var widgets []models.Widget
	dao.DB.Find(&widgets)
	return widgets
}

// AllWidgetsForUser - Returns all widgets for a given user
func (dao *DAO) AllWidgetsForUser(user *models.User) []models.Widget {
	var widgets []models.Widget
	dao.DB.Where("creator_id = ?", user.ID).Find(&widgets)
	return widgets
}

// GetWidgetByID - Get a widget by its ID
func (dao *DAO) GetWidgetByID(ID uint) *models.Widget {
	var widget models.Widget
	if dao.DB.First(&widget, ID).RecordNotFound() {
		return nil
	}
	return &widget
}

// CreateWidget - Create a new widget record.
func (dao *DAO) CreateWidget(widget *models.Widget) *models.Widget {
	dao.DB.Create(&widget)
	return widget
}

// DeleteWidget - Deletes a widget object from the database
func (dao *DAO) DeleteWidget(widget *models.Widget) {
	dao.DB.Delete(widget)
}

// DeleteWidgetByID - Delete a widget record by its primary key
func (dao *DAO) DeleteWidgetByID(id uint) {
	dao.DB.Delete(&models.Widget{ID: id})
}

// UpdateWidget - Updates a widget in the database.
func (dao *DAO) UpdateWidget(w *models.Widget) {
	dao.DB.Model(w).Updates(w)
}
