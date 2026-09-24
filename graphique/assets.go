package main

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Les assets sont intégrés dans l'exécutable : pas besoin de configurer un chemin.
var (
	//go:embed assets/*.png
	assetFS     embed.FS
	assetImages = map[string]*ebiten.Image{}
)

func chargerAssets() {
	noms := []string{
		"floor.png", "wall.png", "player_doc.png",
		"enemy_pumpkin.png", "enemy_zombie.png", "enemy_knight.png",
		"chest.png", "coin.png", "sword.png",
	}
	for _, nom := range noms {
		data, err := assetFS.ReadFile("assets/" + nom)
		if err != nil {
			log.Fatal(err)
		}
		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			log.Fatal(err)
		}
		assetImages[nom] = ebiten.NewImageFromImage(img)
	}
}

func imageAsset(nom string) *ebiten.Image {
	img, ok := assetImages[nom]
	if !ok {
		log.Fatal(fmt.Sprintf("asset introuvable: %s", nom))
	}
	return img
}

func drawFrame(screen *ebiten.Image, sheet *ebiten.Image, frame, frameW, frameH int, x, y, scale float64) {
	frames := sheet.Bounds().Dx() / frameW
	if frames <= 0 {
		return
	}
	frame %= frames
	src := sheet.SubImage(image.Rect(frame*frameW, 0, (frame+1)*frameW, frameH)).(*ebiten.Image)
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x, y)
	screen.DrawImage(src, op)
}

func drawTiled(screen *ebiten.Image, tile *ebiten.Image, tileW, tileH, scale float64) {
	w, h := float64(tileW)*scale, float64(tileH)*scale
	for y := 0.0; y < hauteurEcran; y += h {
		for x := 0.0; x < largeurEcran; x += w {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(scale, scale)
			op.GeoM.Translate(x, y)
			screen.DrawImage(tile, op)
		}
	}
}
