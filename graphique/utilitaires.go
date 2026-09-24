package main
 
import (
	"image/color"
 
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)
 
// ============================================================
// Fonctions utilitaires de dessin, partagées par tous les écrans
// ============================================================
 
// nouvelleImageCouleur crée un carré de couleur unie, en attendant de
// pouvoir le remplacer par un vrai sprite pixel art (fichier .png)
func nouvelleImageCouleur(taille int, c color.RGBA) *ebiten.Image {
	img := ebiten.NewImage(taille, taille)
	img.Fill(c)
	return img
}
 
// seChevauchent teste si deux carrés de même taille se touchent (collision AABB)
func seChevauchent(x1, y1, x2, y2, taille float64) bool {
	return x1 < x2+taille && x1+taille > x2 && y1 < y2+taille && y1+taille > y2
}
 
func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}
 
// dessinerBarre trace une barre de progression (fond gris + remplissage coloré)
func dessinerBarre(screen *ebiten.Image, x, y, largeur, hauteur float32, actuel, max int, couleur color.RGBA) {
	if max <= 0 {
		max = 1
	}
	fraction := float32(actuel) / float32(max)
	if fraction < 0 {
		fraction = 0
	}
	if fraction > 1 {
		fraction = 1
	}
	vector.DrawFilledRect(screen, x, y, largeur, hauteur, color.RGBA{50, 50, 60, 255}, false)
	vector.DrawFilledRect(screen, x, y, largeur*fraction, hauteur, couleur, false)
	vector.StrokeRect(screen, x, y, largeur, hauteur, 1, color.RGBA{200, 200, 200, 255}, false)
}
 
func dessinerBarrePV(screen *ebiten.Image, x, y, largeur, hauteur float32, actuel, max int) {
	fraction := float32(actuel) / float32(maxi(max, 1))
	couleur := color.RGBA{60, 200, 80, 255}
	if fraction < 0.3 {
		couleur = color.RGBA{210, 60, 60, 255}
	} else if fraction < 0.6 {
		couleur = color.RGBA{220, 200, 60, 255}
	}
	dessinerBarre(screen, x, y, largeur, hauteur, actuel, max, couleur)
}
 
func dessinerBarreMana(screen *ebiten.Image, x, y, largeur, hauteur float32, actuel, max int) {
	dessinerBarre(screen, x, y, largeur, hauteur, actuel, max, color.RGBA{70, 140, 220, 255})
}
 
// dessinerPanneau trace un rectangle de fond avec une bordure, pour délimiter
// visuellement une zone d'interface (infos, menu, journal de combat...)
func dessinerPanneau(screen *ebiten.Image, x, y, largeur, hauteur float32) {
	vector.DrawFilledRect(screen, x, y, largeur, hauteur, color.RGBA{25, 25, 35, 230}, false)
	vector.StrokeRect(screen, x, y, largeur, hauteur, 2, color.RGBA{90, 90, 110, 255}, false)
}
 
// dessinerGrille trace une grille discrète en fond d'écran, pour donner un
// effet "carte" pendant l'exploration sans avoir besoin de vraies tuiles
func dessinerGrille(screen *ebiten.Image) {
	pas := float32(40)
	couleur := color.RGBA{45, 45, 60, 255}
	for x := float32(0); x < largeurEcran; x += pas {
		vector.StrokeLine(screen, x, 0, x, hauteurEcran, 1, couleur, false)
	}
	for y := float32(0); y < hauteurEcran; y += pas {
		vector.StrokeLine(screen, 0, y, largeurEcran, y, 1, couleur, false)
	}
}
 