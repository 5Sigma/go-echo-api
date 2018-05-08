package main

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
	port := viper.GetString("port")

	e := echo.New()

	e.Use(middleware.Logger())

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}
	h := Handler{
		DAO: dao.New(),
	}
	endpoints.Setup(e, h)
	vox.Fatal(e.StartServer(s))
}
