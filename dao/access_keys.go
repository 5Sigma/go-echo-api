package dao

import "github.com/5sigma/go-echo-api/models"

// LookupAPIKey - Looks up an access key record by the key. This function
// also preloads the associated user.
func (dao *DAO) LookupAPIKey(key string) *models.APIKey {
	var apiKey models.APIKey
	if dao.DB.Where("key = ?", key).First(&apiKey).RecordNotFound() {
		return nil
	}
	return &apiKey
}

// CreateAPIKeyForUser - Generates a new API key for a user.
func (dao *DAO) CreateAPIKeyForUser(user *models.User) *models.APIKey {
	key := models.APIKey{
		UserID: user.ID,
		Key:    "12312312",
	}
	dao.DB.Create(&key)
	return &key
}
