package main

import (
	"fmt"
	"sort"
	"strings"

	"game-in-go/structures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
)

func nomsRecettes() []string {
	noms := make([]string, 0, len(structures.Recettes))
	for nom := range structures.Recettes {
		noms = append(noms, nom)
	}
	sort.Strings(noms)
	return noms
}

func (g *Game) updateCraft() {
	recettes := nomsRecettes()
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) || inpututil.IsKeyJustPressed(ebiten.Key0) {
		g.etat = EtatAccueil
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.etat = EtatAccueil
		return
	}
	if len(recettes) == 0 {
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.craftIndex = (g.craftIndex - 1 + len(recettes)) % len(recettes)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.craftIndex = (g.craftIndex + 1) % len(recettes)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		nom := recettes[g.craftIndex]
		r := structures.Recettes[nom]
		if g.personnage.InventairePlein() {
			g.craftMessage = "Sac plein : libère une place avant de fabriquer."
		} else if g.personnage.Argent < r.Prix {
			g.craftMessage = fmt.Sprintf("Il faut %d or pour fabriquer %s.", r.Prix, r.NomObjet)
		} else if !g.personnage.PeutFabriquer(r) {
			g.craftMessage = "Ressources insuffisantes : " + ingredientsManquants(*g.personnage, r)
		} else {
			g.personnage.Fabriquer(nom)
			g.craftMessage = "✓ Fabriqué : " + nom
		}
		g.craftMessageTimer = 100
	}
	if g.craftMessageTimer > 0 {
		g.craftMessageTimer--
	}
}

func ingredientsManquants(p structures.Personne, r structures.Recette) string {
	var out []string
	for _, ing := range r.Ingredients {
		n := p.CompteItem(ing.Nom)
		if n < ing.Quantite {
			out = append(out, fmt.Sprintf("%s %d/%d", ing.Nom, n, ing.Quantite))
		}
	}
	return strings.Join(out, ", ")
}

func (g *Game) drawCraft(screen *ebiten.Image) {
	screen.Fill(color.RGBA{10, 10, 17, 255})
	drawTiled(screen, imageAsset("floor.png"), 16, 16, 2)
	dessinerPanneau(screen, 28, 18, 584, 444)
	ebitenutil.DebugPrintAt(screen, "ATELIER DU CAMP", 52, 40)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Or : %d   Sac : %d/%d", g.personnage.Argent, len(g.personnage.Inventaire), g.personnage.TailleMaxInventaireActuelle()), 390, 40)
	ebitenutil.DebugPrintAt(screen, "Fabrique des armes, armures et objets utiles.", 52, 62)

	recettes := nomsRecettes()
	for i, nom := range recettes {
		if i > 8 {
			break
		}
		r := structures.Recettes[nom]
		prefix := "   "
		if i == g.craftIndex {
			prefix = ">> "
		}
		poss := "OK"
		if !g.personnage.PeutFabriquer(r) || g.personnage.Argent < r.Prix {
			poss = "MANQUE"
		}
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s%-23s %2d or  [%s]", prefix, nom, r.Prix, poss), 52, 90+i*27)
	}

	if len(recettes) > 0 {
		r := structures.Recettes[recettes[g.craftIndex]]
		dessinerPanneau(screen, 335, 80, 245, 245)
		ebitenutil.DebugPrintAt(screen, "RECETTE", 355, 103)
		ebitenutil.DebugPrintAt(screen, r.NomObjet, 355, 126)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Coût : %d or", r.Prix), 355, 148)
		ebitenutil.DebugPrintAt(screen, "Ingrédients :", 355, 173)
		y := 194
		for _, ing := range r.Ingredients {
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("- %s : %d/%d", ing.Nom, g.personnage.CompteItem(ing.Nom), ing.Quantite), 355, y)
			y += 20
		}
		if r.Section == "arme" {
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Attaque : +%d", bonusArme(r.NomObjet)), 355, y+5)
		}
		if r.Section == "tete" || r.Section == "torse" || r.Section == "pieds" {
			ebitenutil.DebugPrintAt(screen, "Armure : bonus PV max", 355, y+5)
		}
	}

	ebitenutil.DebugPrintAt(screen, "↑↓ / W,S : choisir     ENTREE : fabriquer     0/Echap : retour", 52, 410)
	ebitenutil.DebugPrintAt(screen, "Astuce : les ressources se trouvent dans les coffres et sur les monstres.", 52, 432)
	if g.craftMessageTimer > 0 {
		dessinerPanneau(screen, 90, 355, 460, 38)
		ebitenutil.DebugPrintAt(screen, g.craftMessage, 105, 371)
	}
}

func bonusArme(nom string) int {
	switch nom {
	case "Dague rouillée":
		return 3
	case "Épée d'acier":
		return 7
	case "Arc du chasseur":
		return 5
	}
	return 0
}
