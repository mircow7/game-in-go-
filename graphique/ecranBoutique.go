package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
)

type articleBoutique struct {
	Nom  string
	Prix int
}

var achatsBoutique = []articleBoutique{
	{"Potion de vie", 25},
	{"Potion de mana", 30},
	{"Potion de poison", 20},
	{"Plume de Corbeau", 12},
	{"Cuir de Sanglier", 15},
	{"Fourrure de Loup", 18},
	{"Peau de Troll", 22},
	{"Fer", 18},
	{"Bois", 10},
	{"Herbe médicinale", 14},
	{"Livre de Sort : Boule de Feu", 120},
}

func prixVente(item string) int {
	switch item {
	case "Potion de vie":
		return 12
	case "Potion de mana":
		return 15
	case "Potion de poison":
		return 10
	case "Plume de Corbeau":
		return 6
	case "Cuir de Sanglier":
		return 8
	case "Fourrure de Loup":
		return 9
	case "Peau de Troll":
		return 11
	case "Fer":
		return 9
	case "Bois":
		return 5
	case "Herbe médicinale":
		return 7
	case "Chapeau de l'aventurier":
		return 45
	case "Tunique de l'aventurier":
		return 60
	case "Bottes de l'aventurier":
		return 50
	case "Livre de Sort : Boule de Feu":
		return 60
	default:
		return 5
	}
}

func (g *Game) updateBoutique() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) || inpututil.IsKeyJustPressed(ebiten.Key0) {
		g.etat = EtatAccueil
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.boutiquePage = 0
		g.boutiqueIndex = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.boutiquePage = 1
		g.boutiqueIndex = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.Key3) {
		g.boutiquePage = 2
		g.boutiqueIndex = 0
	}

	max := len(achatsBoutique)
	if g.boutiquePage == 1 {
		max = len(g.personnage.Inventaire)
	}
	if g.boutiquePage == 2 {
		max = 1
	}
	if max > 0 {
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
			g.boutiqueIndex--
			if g.boutiqueIndex < 0 {
				g.boutiqueIndex = max - 1
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
			g.boutiqueIndex++
			if g.boutiqueIndex >= max {
				g.boutiqueIndex = 0
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.effectuerActionBoutique()
		}
	}
	if g.boutiqueMessageTimer > 0 {
		g.boutiqueMessageTimer--
	}
}

func (g *Game) effectuerActionBoutique() {
	switch g.boutiquePage {
	case 0:
		a := achatsBoutique[g.boutiqueIndex]
		if a.Nom == "Livre de Sort : Boule de Feu" {
			if g.personnage.ConnaitSort("Boule de Feu") {
				g.boutiqueMessage = "Tu connais deja Boule de Feu."
			} else if g.personnage.Argent < a.Prix {
				g.boutiqueMessage = "Pas assez d'or."
			} else {
				g.personnage.Argent -= a.Prix
				g.personnage.SpellBook("Boule de Feu")
				g.boutiqueMessage = "Sort appris : Boule de Feu !"
			}
			g.boutiqueMessageTimer = 90
			return
		}
		if g.personnage.Argent < a.Prix {
			g.boutiqueMessage = "Pas assez d'or."
			g.boutiqueMessageTimer = 90
			return
		}
		if g.personnage.InventairePlein() {
			g.boutiqueMessage = "Sac plein !"
			g.boutiqueMessageTimer = 90
			return
		}
		g.personnage.Argent -= a.Prix
		g.personnage.AddInventory(a.Nom)
		g.boutiqueMessage = fmt.Sprintf("Achat : %s (-%d or)", a.Nom, a.Prix)
	case 1:
		if len(g.personnage.Inventaire) == 0 {
			g.boutiqueMessage = "Ton inventaire est vide."
			g.boutiqueMessageTimer = 90
			return
		}
		item := g.personnage.Inventaire[g.boutiqueIndex]
		prix := prixVente(item)
		g.personnage.RemoveInventory(item)
		g.personnage.Argent += prix
		g.boutiqueMessage = fmt.Sprintf("Vendu : %s (+%d or)", item, prix)
		if g.boutiqueIndex >= len(g.personnage.Inventaire) {
			g.boutiqueIndex = len(g.personnage.Inventaire) - 1
		}
		if g.boutiqueIndex < 0 {
			g.boutiqueIndex = 0
		}
	case 2:
		if g.personnage.AmeliorationsInventaire >= 3 {
			g.boutiqueMessage = "Capacite maximale atteinte."
			g.boutiqueMessageTimer = 90
			return
		}
		prix := 100 * (g.personnage.AmeliorationsInventaire + 1)
		if g.personnage.Argent < prix {
			g.boutiqueMessage = fmt.Sprintf("Il faut %d or.", prix)
			g.boutiqueMessageTimer = 90
			return
		}
		g.personnage.Argent -= prix
		g.personnage.UpgradeInventorySlot()
		g.boutiqueMessage = fmt.Sprintf("Sac agrandi ! %d emplacements.", g.personnage.TailleMaxInventaireActuelle())
	}
	g.boutiqueMessageTimer = 90
}

func (g *Game) drawBoutique(screen *ebiten.Image) {
	screen.Fill(color.RGBA{10, 10, 17, 255})
	drawTiled(screen, imageAsset("floor.png"), 16, 16, 2)
	dessinerPanneau(screen, 35, 22, 570, 440)
	ebitenutil.DebugPrintAt(screen, "BOUTIQUE DU DONJON", 60, 45)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Or : %d", g.personnage.Argent), 470, 45)
	ebitenutil.DebugPrintAt(screen, "[1] Acheter   [2] Vendre   [3] Agrandir le sac", 70, 75)

	switch g.boutiquePage {
	case 0:
		for i, a := range achatsBoutique {
			prefix := "   "
			if i == g.boutiqueIndex {
				prefix = ">> "
			}
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s%d. %-27s %3d or", prefix, i+1, a.Nom, a.Prix), 70, 110+i*30)
		}
		ebitenutil.DebugPrintAt(screen, "Livre de sort : apprend Boule de Feu directement.", 70, 375)
	case 1:
		if len(g.personnage.Inventaire) == 0 {
			ebitenutil.DebugPrintAt(screen, "Aucun objet a vendre.", 70, 120)
		} else {
			for i, item := range g.personnage.Inventaire {
				if i > 11 {
					break
				}
				prefix := "   "
				if i == g.boutiqueIndex {
					prefix = ">> "
				}
				ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s%d. %-30s +%d or", prefix, i+1, item, prixVente(item)), 70, 110+i*30)
			}
		}
		ebitenutil.DebugPrintAt(screen, "Tu peux revendre absolument tous les objets du sac.", 70, 410)
	case 2:
		prochain := 100 * (g.personnage.AmeliorationsInventaire + 1)
		if g.personnage.AmeliorationsInventaire >= 3 {
			ebitenutil.DebugPrintAt(screen, "CAPACITE MAXIMALE : 40 OBJETS", 70, 125)
		} else {
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Capacite actuelle : %d objets", g.personnage.TailleMaxInventaireActuelle()), 70, 125)
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Prochaine amelioration : +10 emplacements pour %d or", prochain), 70, 155)
			ebitenutil.DebugPrintAt(screen, "ENTREE pour acheter l'amelioration", 70, 190)
		}
	}
	ebitenutil.DebugPrintAt(screen, "Fleches / W,S : choisir   ENTREE : valider   0/Echap : retour", 70, 438)
	if g.boutiqueMessageTimer > 0 {
		dessinerPanneau(screen, 110, 385, 420, 38)
		ebitenutil.DebugPrintAt(screen, g.boutiqueMessage, 125, 400)
	}
}
