package router

import (
	"github.com/labstack/echo/v4/middleware"
	"qr-basic-system/internal/infra/api/handler"
	"qr-basic-system/internal/infra/api/router/group"

	"github.com/labstack/echo/v4"
)

type Router struct {
	server  *echo.Echo
	qrGroup group.QR
}

func New(server *echo.Echo, qrGroup group.QR) *Router {
	return &Router{
		server,
		qrGroup,
	}
}

func (r *Router) Init() {
	r.server.Static("/", "client/dist")

	r.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, latency=${latency_human}\n",
	}))

	basePath := r.server.Group("/api/qr-system") //customize your basePath
	basePath.GET("/health", handler.HealthCheck)

	r.qrGroup.Resource(basePath)
}
