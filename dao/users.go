package dao

import "github.com/5sigma/go-echo-api/models"

// AllUsers - Returns all users
func (dao *DAO) AllUsers() []models.User {
	var users []models.User
	dao.DB.Find(&users)
	return users
}

// GetUserByID - Get a user by its ID
func (dao *DAO) GetUserByID(ID uint) *models.User {
	var user models.User
	if dao.DB.First(&user, ID).RecordNotFound() {
		return nil
	}
	return &user
}

// CreateUser - Create a new user record.
func (dao *DAO) CreateUser(user *models.User) *models.User {
	dao.DB.Create(&user)
	return user
}
