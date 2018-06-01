package server

import (
	"fmt"
	"net/http"

	"github.com/5sigma/go-echo-api/dao"
	"github.com/5sigma/go-echo-api/endpoints"
	"github.com/5sigma/vox"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

// StartServer - starts the api server.
func StartServer() {
	var (
		db *dao.DAO
	)
	port := viper.GetString("port")

	e := echo.New()

	e.Use(middleware.Logger())

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}

	if viper.GetBool("memdb") {
		db = dao.NewMemory()
	} else {
		db = dao.New()
	}

	h := endpoints.Handler{
		DB: db,
	}
	endpoints.Setup(e, h)
	vox.Fatal(e.StartServer(server))
}
