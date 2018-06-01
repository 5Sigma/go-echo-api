package dao

import (
	"testing"

	"github.com/5sigma/go-echo-api/models"
)

func TestAllUsers(t *testing.T) {
	dao := NewMemory()
	dao.CreateUser(&models.User{})
	dao.CreateUser(&models.User{})
	dao.CreateUser(&models.User{})
	users := dao.AllUsers()
	if len(users) != 3 {
		t.Errorf("Users were not returned. Count is %d", len(users))
	}
}

func TestGetUserByID(t *testing.T) {
	dao := NewMemory()
	user := dao.CreateUser(&models.User{})
	user = dao.GetUserByID(user.ID)
	if user == nil {
		t.Errorf("No user found for ID %d", user.ID)
	}
	user = dao.GetUserByID(8)
	if user != nil {
		t.Error("User returned with bad ID")
	}
}

func TestCreateUser(t *testing.T) {
	dao := NewMemory()
	user := dao.CreateUser(&models.User{})
	if user.ID == 0 {
		t.Error("ID did not come back after creating a user")
	}
}
