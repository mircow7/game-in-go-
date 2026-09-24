package main

import (
	"fmt"
	"game-in-go/structures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
)

func (g *Game) updateInventaire() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) || inpututil.IsKeyJustPressed(ebiten.Key0) {
		g.etat = EtatAccueil
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.inventaireIndex--
		if g.inventaireIndex < 0 {
			g.inventaireIndex = len(g.personnage.Inventaire) - 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.inventaireIndex++
		if g.inventaireIndex >= len(g.personnage.Inventaire) {
			g.inventaireIndex = 0
		}
	}
	if len(g.personnage.Inventaire) == 0 {
		g.inventaireIndex = 0
		return
	}
	if g.inventaireIndex < 0 {
		g.inventaireIndex = 0
	}
	if g.inventaireIndex >= len(g.personnage.Inventaire) {
		g.inventaireIndex = len(g.personnage.Inventaire) - 1
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		item := g.personnage.Inventaire[g.inventaireIndex]
		switch item {
		case "Potion de vie":
			g.personnage.TakePot()
			g.inventaireMessage = "Potion de vie utilisee."
		case "Potion de mana":
			g.personnage.TakeManaPot()
			g.inventaireMessage = "Potion de mana utilisee."
		case "Potion de poison":
			g.inventaireMessage = "Garde cette potion pour un ennemi : elle est dangereuse pour toi !"
		case "Potion supérieure":
			if g.personnage.RemoveInventory(item) {
				g.personnage.PointsDeVieActuels += 100
				if g.personnage.PointsDeVieActuels > g.personnage.PointsDeVieMaximum {
					g.personnage.PointsDeVieActuels = g.personnage.PointsDeVieMaximum
				}
				g.inventaireMessage = "Potion supérieure utilisée : +100 PV."
			}
		case "Elixir de mana":
			if g.personnage.RemoveInventory(item) {
				g.personnage.Mana += 50
				if g.personnage.Mana > g.personnage.ManaMax {
					g.personnage.Mana = g.personnage.ManaMax
				}
				g.inventaireMessage = "Elixir de mana utilisé : +50 mana."
			}
		default:
			if r, ok := structures.Recettes[item]; ok {
				g.personnage.Equiper(item, r.Section)
				g.inventaireMessage = item + " equipe."
				if g.inventaireIndex >= len(g.personnage.Inventaire) {
					g.inventaireIndex = len(g.personnage.Inventaire) - 1
				}
			} else {
				g.inventaireMessage = "Objet non utilisable ici."
			}
		}
		g.inventaireMessageTimer = 90
	}
	if g.inventaireMessageTimer > 0 {
		g.inventaireMessageTimer--
	}
}

func (g *Game) drawInventaire(screen *ebiten.Image) {
	screen.Fill(color.RGBA{10, 10, 17, 255})
	drawTiled(screen, imageAsset("floor.png"), 16, 16, 2)
	dessinerPanneau(screen, 45, 25, 550, 430)
	ebitenutil.DebugPrintAt(screen, "INVENTAIRE", 70, 48)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Or : %d    Sac : %d/%d", g.personnage.Argent, len(g.personnage.Inventaire), g.personnage.TailleMaxInventaireActuelle()), 70, 70)
	dessinerBarrePV(screen, 70, 90, 180, 12, g.personnage.PointsDeVieActuels, g.personnage.PointsDeVieMaximum)
	dessinerBarreMana(screen, 270, 90, 180, 10, g.personnage.Mana, g.personnage.ManaMax)

	if len(g.personnage.Inventaire) == 0 {
		ebitenutil.DebugPrintAt(screen, "Ton sac est vide.", 80, 135)
	} else {
		for i, item := range g.personnage.Inventaire {
			if i > 13 {
				break
			}
			prefix := "   "
			if i == g.inventaireIndex {
				prefix = ">> "
			}
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s%d. %s", prefix, i+1, item), 80, 125+i*22)
		}
	}
	ebitenutil.DebugPrintAt(screen, "ENTREE : utiliser/equiper    0/Echap : retour", 80, 420)
	ebitenutil.DebugPrintAt(screen, "Les potions soignent ici sans lancer de combat.", 80, 438)
	if g.inventaireMessageTimer > 0 {
		dessinerPanneau(screen, 120, 350, 400, 42)
		ebitenutil.DebugPrintAt(screen, g.inventaireMessage, 135, 367)
	}
}
