package main

import (
	"fmt"
	"log"

	"qr-basic-system/cmd/providers"
	"qr-basic-system/config"
	"qr-basic-system/internal/infra/api/router"
	"github.com/labstack/echo/v4"
)

var (
	serverHost = config.Environments().ServerHost
	serverPort = config.Environments().ServerPort
)

func main() {
	container := providers.BuildContainer()

	err := container.Invoke(func(router *router.Router, server *echo.Echo) {
		router.Init()
		server.Logger.Fatal(server.Start(fmt.Sprintf("%s:%d", serverHost, serverPort)))
	})

	if err != nil {
		log.Panic(err)
	}
}
