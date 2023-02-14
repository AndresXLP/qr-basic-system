package app

import (
	"bytes"
	"fmt"
	"image/png"
	"os"

	"github.com/goombaio/namegenerator"
	"github.com/skip2/go-qrcode"
)

type QR interface {
	GenerateQR(content string) (string, error)
}

type qr struct {
	nameGenerator namegenerator.Generator
}

func NewQr(generator namegenerator.Generator) QR {
	return &qr{generator}
}

func (q *qr) GenerateQR(content string) (string, error) {
	qrByte, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}

	img, err := png.Decode(bytes.NewBuffer(qrByte))
	if err != nil {
		return "", err
	}

	name := fmt.Sprintf("tmp/%s.png", q.nameGenerator.Generate())
	file, err := os.Create(name)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if err = png.Encode(file, img); err != nil {
		return "", err
	}

	return name, nil
}
