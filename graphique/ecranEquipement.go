package main

import (
	"fmt"
	"game-in-go/structures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
)

func (g *Game) objetsEquipables() []string {
	var r []string
	for _, item := range g.personnage.Inventaire {
		if _, ok := structures.Recettes[item]; ok {
			r = append(r, item)
		}
	}
	return r
}

func (g *Game) updateEquipement() {
	items := g.objetsEquipables()
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) || inpututil.IsKeyJustPressed(ebiten.Key0) {
		g.etat = EtatAccueil
		return
	}
	if len(items) == 0 {
		g.equipementIndex = 0
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.equipementIndex--
		if g.equipementIndex < 0 {
			g.equipementIndex = len(items) - 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.equipementIndex++
		if g.equipementIndex >= len(items) {
			g.equipementIndex = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		item := items[g.equipementIndex]
		recette := structures.Recettes[item]
		g.personnage.Equiper(item, recette.Section)
		g.equipementMessage = item + " equipe !"
		g.equipementMessageTimer = 90
		if g.equipementIndex >= len(g.objetsEquipables()) {
			g.equipementIndex = len(g.objetsEquipables()) - 1
		}
		if g.equipementIndex < 0 {
			g.equipementIndex = 0
		}
	}
	if g.equipementMessageTimer > 0 {
		g.equipementMessageTimer--
	}
}

func (g *Game) drawEquipement(screen *ebiten.Image) {
	screen.Fill(color.RGBA{10, 10, 17, 255})
	drawTiled(screen, imageAsset("floor.png"), 16, 16, 2)
	dessinerPanneau(screen, 35, 25, 570, 430)
	ebitenutil.DebugPrintAt(screen, "EQUIPEMENT", 65, 50)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("PV max : %d", g.personnage.PointsDeVieMaximum), 65, 78)
	ebitenutil.DebugPrintAt(screen, "TETE  : "+videSiVide(g.personnage.Equipement.Tete), 65, 110)
	ebitenutil.DebugPrintAt(screen, "TORSE : "+videSiVide(g.personnage.Equipement.Torse), 65, 132)
	ebitenutil.DebugPrintAt(screen, "PIEDS : "+videSiVide(g.personnage.Equipement.Pieds), 65, 154)
	ebitenutil.DebugPrintAt(screen, "ARME  : "+videSiVide(g.personnage.Equipement.Arme)+fmt.Sprintf(" (+%d attaque)", g.personnage.BonusAttaqueEquipement()), 65, 176)
	ebitenutil.DebugPrintAt(screen, "Objets equipables dans le sac :", 65, 215)
	items := g.objetsEquipables()
	if len(items) == 0 {
		ebitenutil.DebugPrintAt(screen, "Aucun equipement. Va dans Crafting pour fabriquer du stuff.", 65, 245)
	} else {
		for i, item := range items {
			prefix := "   "
			if i == g.equipementIndex {
				prefix = ">> "
			}
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s%d. %s (+PV)", prefix, i+1, item), 80, 240+i*26)
		}
	}
	ebitenutil.DebugPrintAt(screen, "ENTREE : equiper   0/Echap : retour", 65, 420)
	if g.equipementMessageTimer > 0 {
		dessinerPanneau(screen, 180, 365, 280, 38)
		ebitenutil.DebugPrintAt(screen, g.equipementMessage, 200, 380)
	}
}

func videSiVide(s string) string {
	if s == "" {
		return "Aucun"
	}
	return s
}
