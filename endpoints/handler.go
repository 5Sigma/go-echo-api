package endpoints

import "github.com/5sigma/go-echo-api/dao"

// Handler - API Request handler
type Handler struct {
	DB *dao.DAO
}
