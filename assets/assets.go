package assets

import (
	"bytes"
	"embed"
	"image"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var embeddedAssets embed.FS

func LoadImageFromAssets(filename string) *ebiten.Image {
	imgData, err := embeddedAssets.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read embedded file: %s, error: %v", filename, err)
	}

	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		log.Fatalf("Failed to decode image: %s, error: %v", filename, err)
	}

	ebitenImage := ebiten.NewImageFromImage(img)
	return ebitenImage
}
