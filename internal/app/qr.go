package app

import (
	"github.com/goombaio/namegenerator"
	"github.com/skip2/go-qrcode"
)

type QR interface {
	GenerateQR(content string) ([]byte, error)
}

type qr struct {
	nameGenerator namegenerator.Generator
}

func NewQr(generator namegenerator.Generator) QR {
	return &qr{generator}
}

func (q *qr) GenerateQR(content string) ([]byte, error) {
	qrByte, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	return qrByte, nil
}
