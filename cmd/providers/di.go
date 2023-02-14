package providers

import (
	"time"

	"github.com/goombaio/namegenerator"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
	"qr-basic-system/internal/app"
	"qr-basic-system/internal/infra/api/handler"
	"qr-basic-system/internal/infra/api/router"
	"qr-basic-system/internal/infra/api/router/group"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})
	
	_ = Container.Provide(func() namegenerator.Generator {
		return namegenerator.NewNameGenerator(time.Now().Local().UTC().UnixNano())
	})

	_ = Container.Provide(router.New)

	_ = Container.Provide(group.NewQr)

	_ = Container.Provide(handler.NewQr)

	_ = Container.Provide(app.NewQr)

	return Container
}
