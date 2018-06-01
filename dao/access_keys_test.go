package dao

import (
	"testing"

	"github.com/5sigma/go-echo-api/models"
)

func TestLookupAPIKey(t *testing.T) {
	var (
		res *models.APIKey
	)
	dao := NewMemory()
	user := dao.CreateUser(&models.User{})
	key := dao.CreateAPIKeyForUser(user)
	res = dao.LookupAPIKey("badkey")
	if res != nil {
		t.Errorf("Non nil response for bad key. Key value is '%s'", res.Key)
	}
	res = dao.LookupAPIKey(key.Key)
	if res == nil {
		t.Error("Nil repsonse for good key value")
	}
}
