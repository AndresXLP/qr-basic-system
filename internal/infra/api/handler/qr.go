package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"qr-basic-system/internal/app"
	"qr-basic-system/internal/domain/dto"
)

type QR interface {
	Generate(cntx echo.Context) error
}

type qr struct {
	service app.QR
}

func NewQr(service app.QR) QR {
	return &qr{
		service,
	}
}

func (q *qr) Generate(cntx echo.Context) error {
	var requestBody dto.QR
	if err := cntx.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := requestBody.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	name, err := q.service.GenerateQR(requestBody.Content)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	//defer os.Remove(name)

	return cntx.JSON(http.StatusCreated, name)
}
