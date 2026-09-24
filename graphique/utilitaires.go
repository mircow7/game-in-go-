package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

func seChevauchent(x1, y1, x2, y2, taille float64) bool {
	return x1 < x2+taille && x1+taille > x2 && y1 < y2+taille && y1+taille > y2
}
func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dessinerBarre(screen *ebiten.Image, x, y, largeur, hauteur float32, actuel, max int, couleur color.RGBA) {
	if max <= 0 {
		max = 1
	}
	f := float32(actuel) / float32(max)
	if f < 0 {
		f = 0
	}
	if f > 1 {
		f = 1
	}
	vector.DrawFilledRect(screen, x, y, largeur, hauteur, color.RGBA{35, 35, 45, 255}, false)
	vector.DrawFilledRect(screen, x, y, largeur*f, hauteur, couleur, false)
	vector.StrokeRect(screen, x, y, largeur, hauteur, 1, color.RGBA{180, 180, 195, 255}, false)
}
func dessinerBarrePV(screen *ebiten.Image, x, y, largeur, hauteur float32, actuel, max int) {
	f := float32(actuel) / float32(maxi(max, 1))
	c := color.RGBA{70, 210, 95, 255}
	if f < .3 {
		c = color.RGBA{220, 65, 65, 255}
	} else if f < .6 {
		c = color.RGBA{225, 195, 55, 255}
	}
	dessinerBarre(screen, x, y, largeur, hauteur, actuel, max, c)
}
func dessinerBarreMana(screen *ebiten.Image, x, y, largeur, hauteur float32, actuel, max int) {
	dessinerBarre(screen, x, y, largeur, hauteur, actuel, max, color.RGBA{65, 135, 235, 255})
}
func dessinerBarreXP(screen *ebiten.Image, x, y, largeur, hauteur float32, actuel, max int) {
	dessinerBarre(screen, x, y, largeur, hauteur, actuel, max, color.RGBA{190, 90, 235, 255})
}
func dessinerPanneau(screen *ebiten.Image, x, y, largeur, hauteur float32) {
	vector.DrawFilledRect(screen, x, y, largeur, hauteur, color.RGBA{18, 18, 28, 235}, false)
	vector.StrokeRect(screen, x, y, largeur, hauteur, 2, color.RGBA{105, 105, 130, 255}, false)
}
func dessinerGrille(screen *ebiten.Image) {
	// très légère grille décorative
	for x := float32(0); x < largeurEcran; x += 32 {
		vector.StrokeLine(screen, x, 0, x, hauteurEcran, 1, color.RGBA{40, 40, 50, 80}, false)
	}
	for y := float32(0); y < hauteurEcran; y += 32 {
		vector.StrokeLine(screen, 0, y, largeurEcran, y, 1, color.RGBA{40, 40, 50, 80}, false)
	}
}
