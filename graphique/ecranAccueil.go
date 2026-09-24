package main

import (
	"fmt"
	"os"

	"game-in-go/structures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

var accueilOptions = []string{
	"Explorer le donjon",
	"Inventaire",
	"Boutique",
	"Equipement",
	"Crafting",
	"Sauvegarder",
	"Quitter",
}

func (g *Game) updateAccueil() {
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.craftIndex = 0
		g.etat = EtatCraft
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.accueilIndex--
		if g.accueilIndex < 0 {
			g.accueilIndex = len(accueilOptions) - 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.accueilIndex++
		if g.accueilIndex >= len(accueilOptions) {
			g.accueilIndex = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.etat = EtatTitre
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		switch g.accueilIndex {
		case 0:
			g.etat = EtatExploration
		case 1:
			g.etat = EtatInventaire
		case 2:
			g.boutiquePage = 0
			g.boutiqueIndex = 0
			g.etat = EtatBoutique
		case 3:
			g.equipementIndex = 0
			g.etat = EtatEquipement
		case 4:
			g.craftIndex = 0
			g.etat = EtatCraft
		case 5:
			if err := structures.Sauvegarder(*g.personnage); err != nil {
				g.accueilMessage = "Erreur de sauvegarde"
			} else {
				g.accueilMessage = "Partie sauvegardee !"
			}
			g.accueilMessageTimer = 90
		case 6:
			os.Exit(0)
		}
	}
	if g.accueilMessageTimer > 0 {
		g.accueilMessageTimer--
	}
}

func (g *Game) drawAccueil(screen *ebiten.Image) {
	screen.Fill(color.RGBA{9, 9, 16, 255})
	drawTiled(screen, imageAsset("floor.png"), 16, 16, 2)
	vector.DrawFilledRect(screen, 0, 0, largeurEcran, hauteurEcran, color.RGBA{5, 5, 12, 175}, false)

	// Decoration: hero, coffre et ennemi.
	drawFrame(screen, imageAsset("player_doc.png"), g.animTick/8, 16, 23, 58, 170, 4)
	drawFrame(screen, imageAsset("enemy_knight.png"), g.animTick/8, 16, 28, 540, 160, 4)
	drawFrame(screen, imageAsset("chest.png"), g.animTick/10, 16, 16, 74, 365, 3)
	drawFrame(screen, imageAsset("coin.png"), g.animTick/8, 6, 7, 545, 375, 4)

	dessinerPanneau(screen, 125, 42, 390, 390)
	ebitenutil.DebugPrintAt(screen, "PROJET RED", 260, 65)
	ebitenutil.DebugPrintAt(screen, "CAMP DE BASE", 240, 88)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s  |  Niv.%d", g.personnage.Nom, g.personnage.Niveau), 220, 120)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("PV %d/%d", g.personnage.PointsDeVieActuels, g.personnage.PointsDeVieMaximum), 180, 145)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Mana %d/%d", g.personnage.Mana, g.personnage.ManaMax), 350, 145)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Or : %d", g.personnage.Argent), 180, 168)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Sac : %d/%d", len(g.personnage.Inventaire), g.personnage.TailleMaxInventaireActuelle()), 350, 168)

	for i, option := range accueilOptions {
		prefix := "   "
		if i == g.accueilIndex {
			prefix = ">> "
		}
		ebitenutil.DebugPrintAt(screen, prefix+option, 205, 205+i*32)
	}

	ebitenutil.DebugPrintAt(screen, "Fleches / W,S : choisir    ENTREE : valider", 165, 425)
	ebitenutil.DebugPrintAt(screen, "C : crafting rapide   H : revenir ici depuis le donjon", 175, 443)
	if g.accueilMessageTimer > 0 {
		dessinerPanneau(screen, 190, 15, 260, 30)
		ebitenutil.DebugPrintAt(screen, g.accueilMessage, 215, 25)
	}
	g.animTick++
}
