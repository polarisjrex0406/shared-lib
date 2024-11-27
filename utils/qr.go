package utils

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func GenerateQRCode(data, path string) error {
	// Generate QR code
	qrCode, err := qr.Encode(data, qr.L, qr.Unicode)
	if err != nil {
		return err
	}

	// Scale the QR code to a larger size (e.g., 128x128 pixels)
	qrCode, err = barcode.Scale(qrCode, 128, 128) // Adjust width and height as needed
	if err != nil {
		return err
	}

	// Save QR code as an image
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	png.Encode(file, qrCode)

	return nil
}
