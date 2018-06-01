package dao

import (
	"testing"

	"github.com/5sigma/go-echo-api/models"
)

func TestAllWidgets(t *testing.T) {
	dao := NewMemory()
	dao.CreateWidget(&models.Widget{})
	dao.CreateWidget(&models.Widget{})
	dao.CreateWidget(&models.Widget{})
	widgets := dao.AllWidgets()
	if len(widgets) != 3 {
		t.Errorf("Widgets were not returned. Count is %d", len(widgets))
	}
}

func TestGetWidgetByID(t *testing.T) {
	dao := NewMemory()
	widget := dao.CreateWidget(&models.Widget{})
	widget = dao.GetWidgetByID(widget.ID)
	if widget == nil {
		t.Errorf("No widget found for ID %d", widget.ID)
	}
	widget = dao.GetWidgetByID(8)
	if widget != nil {
		t.Error("Widget returned with bad ID")
	}
}

func TestCreateWidget(t *testing.T) {
	dao := NewMemory()
	widget := &models.Widget{}
	widget = dao.CreateWidget(widget)
	if widget.ID == 0 {
		t.Error("ID did not come back after creating a widget")
	}
}

func TestDeleteWidget(t *testing.T) {
	dao := NewMemory()
	widget := dao.CreateWidget(&models.Widget{})
	dao.DeleteWidget(widget)
	widgets := dao.AllWidgets()
	if len(widgets) != 0 {
		t.Errorf("Widgets were not deleted: Count is %d", len(widgets))
	}
}

func TestUpdateWidget(t *testing.T) {
	dao := NewMemory()
	widget := dao.CreateWidget(&models.Widget{})
	widget.Name = "test"
	dao.UpdateWidget(widget)
	widget = dao.GetWidgetByID(widget.ID)
	if widget.Name != "test" {
		t.Errorf("Widget name not updated: '%s'", widget.Name)
	}
}

func TestGetAllWidgetsForUser(t *testing.T) {
	dao := NewMemory()
	user := dao.CreateUser(&models.User{})
	dao.CreateWidget(&models.Widget{CreatorID: user.ID})
	dao.CreateWidget(&models.Widget{})
	dao.CreateWidget(&models.Widget{})
	widgets := dao.AllWidgetsForUser(user)
	if len(widgets) != 1 {
		t.Errorf("Widgets were not returned. Count is %d", len(widgets))
	}
}

func TestDeleteWidgetByID(t *testing.T) {
	dao := NewMemory()
	user := dao.CreateUser(&models.User{})
	widget := dao.CreateWidget(&models.Widget{CreatorID: user.ID})
	dao.DeleteWidgetByID(widget.ID)
	widgets := dao.AllWidgetsForUser(user)
	if len(widgets) != 0 {
		t.Errorf("Widgets were not returned. Count is %d", len(widgets))
	}
}
