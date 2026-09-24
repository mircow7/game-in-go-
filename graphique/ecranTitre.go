package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

func (g *Game) updateTitre() {
	g.animTick++
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.etat = EtatAccueil
	}
}

func (g *Game) drawTitre(screen *ebiten.Image) {
	screen.Fill(color.RGBA{10, 10, 17, 255})
	drawTiled(screen, imageAsset("floor.png"), 16, 16, 2)
	vector.DrawFilledRect(screen, 0, 0, 640, 480, color.RGBA{8, 8, 15, 190}, false)

	// Héros et monstre animés.
	drawFrame(screen, imageAsset("player_doc.png"), g.animTick/10, 16, 23, 105, 125, 4)
	drawFrame(screen, imageAsset("enemy_pumpkin.png"), g.animTick/10, 16, 23, 485, 125, 4)

	dessinerPanneau(screen, 75, 70, 490, 315)
	ebitenutil.DebugPrintAt(screen, "PROJECT RED", 250, 95)
	ebitenutil.DebugPrintAt(screen, "DUNGEON EDITION", 218, 118)
	ebitenutil.DebugPrintAt(screen, "Une nouvelle aventure t'attend...", 180, 165)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s  |  Niveau %d", g.personnage.Nom, g.personnage.Niveau), 220, 205)
	ebitenutil.DebugPrintAt(screen, "ENTREE  :  ouvrir le camp", 235, 250)
	ebitenutil.DebugPrintAt(screen, "Avant le donjon : boutique, inventaire, equipement", 200, 278)
	ebitenutil.DebugPrintAt(screen, "Echap : pause   |   1/2/3 : combat", 185, 300)
	ebitenutil.DebugPrintAt(screen, "Trouve les coffres, récupère l'or", 190, 340)
	ebitenutil.DebugPrintAt(screen, "et affronte les créatures du donjon !", 170, 358)
	g.animTick++
}
