package router

import (
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
	basePath := r.server.Group("/api/qr-system") //customize your basePath
	basePath.GET("/health", handler.HealthCheck)

	r.qrGroup.Resource(basePath)
}
