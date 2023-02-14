package group

import (
	"github.com/labstack/echo/v4"
	"qr-basic-system/internal/infra/api/handler"
)

type QR interface {
	Resource(g *echo.Group)
}

type qr struct {
	handler handler.QR
}

func NewQr(hand handler.QR) QR {
	return &qr{hand}
}

func (q *qr) Resource(g *echo.Group) {
	groupPath := g.Group("/generate")
	groupPath.POST("", q.handler.Generate)
}
