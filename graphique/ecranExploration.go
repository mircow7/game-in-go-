package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand"

	"game-in-go/structures"
)

func (g *Game) positionnerMonstreAleatoirement() {
	positions := [][2]float64{{180, 110}, {390, 115}, {210, 290}, {470, 315}, {105, 340}}
	p := positions[rand.Intn(len(positions))]
	g.monstreX, g.monstreY = p[0], p[1]
	g.monstreActif = true
	g.monstreExploration = randomMonsterForMap()
}

func randomMonsterForMap() structures.Monster {
	// Cette fonction est remplacée dans le fichier compilé par l'import ci-dessous.
	return structures.ChoisirMonsterAleatoire()
}

func (g *Game) updateExploration() {
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		g.etat = EtatAccueil
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.etat = EtatPause
		g.pauseIndex = 0
		return
	}
	if !g.monstreActif {
		g.framesAvantRespawn--
		if g.framesAvantRespawn <= 0 {
			g.positionnerMonstreAleatoirement()
		}
	}
	if g.coinActif && seChevauchent(g.joueurX, g.joueurY, g.coinX, g.coinY, 26) {
		g.personnage.Argent += 5
		g.coinActif = false
		g.combatPopup = "+5 OR"
		g.combatPopupTimer = 90
	}
	if g.coffreActif && seChevauchent(g.joueurX, g.joueurY, g.coffreX, g.coffreY, 30) {
		g.coffreActif = false
		g.personnage.Argent += 20
		g.personnage.AddInventory("Potion de vie")
		g.combatPopup = "COFFRE : +20 OR + POTION"
		g.combatPopupTimer = 120
	}
	vitesse := 3.0
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.joueurX -= vitesse
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.joueurX += vitesse
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.joueurY -= vitesse
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.joueurY += vitesse
	}
	if g.joueurX < 34 {
		g.joueurX = 34
	}
	if g.joueurY < 42 {
		g.joueurY = 42
	}
	if g.joueurX > 574 {
		g.joueurX = 574
	}
	if g.joueurY > 405 {
		g.joueurY = 405
	}

	g.animTick++
	if g.animTick%8 == 0 {
		g.frameAnim = (g.frameAnim + 1) % 4
	}
	if g.coinTick++; g.coinTick%10 == 0 && g.combatPopupTimer > 0 {
		g.combatPopupTimer--
	}
	if g.combatPopupTimer > 0 {
		g.combatPopupTimer--
	}

	if g.monstreActif && seChevauchent(g.joueurX, g.joueurY, g.monstreX, g.monstreY, 28) {
		g.demarrerCombat()
	}
}

func (g *Game) drawDungeonBackground(screen *ebiten.Image) {
	drawTiled(screen, imageAsset("floor.png"), 16, 16, 2)
	// Bordure en pierre.
	wall := imageAsset("wall.png")
	for x := 32.0; x < 608; x += 32 {
		for _, y := range []float64{0, 416} {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(2, 2)
			op.GeoM.Translate(x, y)
			screen.DrawImage(wall, op)
		}
	}
	for y := 32.0; y < 416; y += 32 {
		for _, x := range []float64{0, 608} {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(2, 2)
			op.GeoM.Translate(x, y)
			screen.DrawImage(wall, op)
		}
	}
	// Quelques blocs décoratifs pour donner l'impression d'une vraie salle.
	vector.DrawFilledRect(screen, 64, 64, 32, 320, color.RGBA{10, 10, 18, 70}, false)
	vector.DrawFilledRect(screen, 544, 64, 32, 320, color.RGBA{10, 10, 18, 70}, false)
}

func (g *Game) drawExploration(screen *ebiten.Image) {
	g.drawDungeonBackground(screen)
	// Coffre
	if g.coffreActif {
		drawFrame(screen, imageAsset("chest.png"), g.frameAnim, 16, 16, g.coffreX, g.coffreY, 2)
	}
	// Pièce
	if g.coinActif {
		drawFrame(screen, imageAsset("coin.png"), g.coinTick/8, 6, 7, g.coinX, g.coinY, 3)
	}
	// Ennemi
	if g.monstreActif {
		sprite := "enemy_pumpkin.png"
		switch g.monstreExploration.SpriteID {
		case "zombie":
			sprite = "enemy_zombie.png"
		case "knight":
			sprite = "enemy_knight.png"
		}
		w, h := 16, 23
		if sprite == "enemy_zombie.png" {
			w, h = 16, 16
		}
		if sprite == "enemy_knight.png" {
			w, h = 16, 28
		}
		drawFrame(screen, imageAsset(sprite), g.frameAnim, w, h, g.monstreX, g.monstreY-14, 2)
		// petite barre de vie
		dessinerBarrePV(screen, float32(g.monstreX-4), float32(g.monstreY-18), 40, 5, g.monstreExploration.PointsDeVieActuels, g.monstreExploration.PointsDeVieMaximum)
		ebitenutil.DebugPrintAt(screen, g.monstreExploration.Nom, int(g.monstreX-12), int(g.monstreY+38))
	}
	// Joueur
	drawFrame(screen, imageAsset("player_doc.png"), g.frameAnim, 16, 23, g.joueurX, g.joueurY-14, 2)

	// Vignette légère
	vector.StrokeRect(screen, 6, 6, 628, 424, 2, color.RGBA{120, 120, 145, 255}, false)

	// HUD
	dessinerPanneau(screen, 0, 430, 640, 50)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("DOC  | Niv.%d | PV %d/%d | Or %d", g.personnage.Niveau, g.personnage.PointsDeVieActuels, g.personnage.PointsDeVieMaximum, g.personnage.Argent), 12, 438)
	dessinerBarreXP(screen, 280, 439, 145, 8, g.personnage.ExperienceActuelle, g.personnage.ExperienceMax)
	ebitenutil.DebugPrintAt(screen, "Fleches/WASD | H : camp | Echap : pause", 440, 438)
	if g.combatPopupTimer > 0 {
		dessinerPanneau(screen, 155, 185, 330, 55)
		ebitenutil.DebugPrintAt(screen, g.combatPopup, 180, 208)
	}
}
