package dao

import (
	"github.com/5sigma/go-echo-api/models"
	"github.com/jinzhu/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.APIKey{},
		&models.Widget{},
	)
}
