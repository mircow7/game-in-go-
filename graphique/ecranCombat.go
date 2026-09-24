package main

import (
	"fmt"
	"game-in-go/structures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type PhaseCombat int

const (
	PhaseMenu PhaseCombat = iota
	PhaseSorts
	PhaseInventaire
	PhaseTourMonstre
	PhaseVictoire
	PhaseDefaite
)

var touchesNombres = []ebiten.Key{ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.Key5, ebiten.Key6, ebiten.Key7, ebiten.Key8, ebiten.Key9}

func (g *Game) ajouterMessage(msg string) {
	g.combatMessages = append(g.combatMessages, msg)
	if len(g.combatMessages) > 6 {
		g.combatMessages = g.combatMessages[len(g.combatMessages)-6:]
	}
}

func (g *Game) demarrerCombat() {
	g.combatMonstre = g.monstreExploration
	if g.combatMonstre.Nom == "" {
		g.combatMonstre = structures.ChoisirMonsterAleatoire()
	}
	g.etat = EtatCombat
	g.combatMessages = nil
	g.combatTourMonstre = 1
	g.combatPhase = PhaseMenu
	g.flashJoueur = 0
	g.flashMonstre = 0
	g.combatAnimTick = 0
	g.combatFrame = 0
	g.ajouterMessage("⚔ " + g.combatMonstre.Nom + " apparaît !")
	premier := structures.DeterminerPremierJoueur(g.personnage, &g.combatMonstre)
	if premier != "joueur" {
		g.ajouterMessage("L'ennemi est plus rapide !")
		g.combatPhase = PhaseTourMonstre
		g.combatMinuteur = 40
	} else {
		g.ajouterMessage("Tu prends l'initiative !")
	}
}

func (g *Game) apresActionJoueur() {
	if g.combatMonstre.EstMort() {
		g.ajouterMessage("✦ " + g.combatMonstre.Nom + " est vaincu !")
		g.personnage.GainExperience(g.combatMonstre.ExperienceOffert)
		gain := g.combatMonstre.OrRecompense
		g.personnage.Argent += gain
		g.combatPopup = fmt.Sprintf("+%d OR  |  +%d XP", gain, g.combatMonstre.ExperienceOffert)
		g.combatPopupTimer = 100
		if objet, obtenu := structures.TirerButin(g.combatMonstre); obtenu {
			if g.personnage.InventairePlein() {
				g.ajouterMessage("Butin trouvé mais inventaire plein !")
			} else {
				g.personnage.AddInventory(objet)
				g.ajouterMessage("Butin : " + objet)
			}
		}
		g.combatPhase = PhaseVictoire
		g.monstreActif = false
		g.framesAvantRespawn = 120
		return
	}
	g.combatPhase = PhaseTourMonstre
	g.combatMinuteur = 40
}

func (g *Game) utiliserObjetCombat(item string) {
	switch item {
	case "Potion de vie":
		g.personnage.TakePot()
		g.ajouterMessage("Potion de vie utilisée.")
	case "Potion de poison":
		g.personnage.PoisonPot()
		g.ajouterMessage("Aïe... poison !")
	case "Potion de mana":
		g.personnage.TakeManaPot()
		g.ajouterMessage("Potion de mana utilisée.")
	case "Potion supérieure":
		if g.personnage.RemoveInventory(item) {
			g.personnage.PointsDeVieActuels += 100
			if g.personnage.PointsDeVieActuels > g.personnage.PointsDeVieMaximum {
				g.personnage.PointsDeVieActuels = g.personnage.PointsDeVieMaximum
			}
			g.ajouterMessage("Potion supérieure : +100 PV.")
		}
	case "Elixir de mana":
		if g.personnage.RemoveInventory(item) {
			g.personnage.Mana += 50
			if g.personnage.Mana > g.personnage.ManaMax {
				g.personnage.Mana = g.personnage.ManaMax
			}
			g.ajouterMessage("Elixir : +50 mana.")
		}
	case "AK-47":
		structures.UtiliserAK47(g.personnage, &g.combatMonstre)
		g.flashMonstre = 10
		g.ajouterMessage("Rafale !")
	default:
		g.ajouterMessage("Cet objet ne s'utilise pas en combat.")
	}
}

func (g *Game) updateCombat() {
	if g.flashJoueur > 0 {
		g.flashJoueur--
	}
	if g.flashMonstre > 0 {
		g.flashMonstre--
	}
	g.combatAnimTick++
	if g.combatAnimTick%8 == 0 {
		g.combatFrame = (g.combatFrame + 1) % 4
	}
	if g.combatPopupTimer > 0 {
		g.combatPopupTimer--
	}
	switch g.combatPhase {
	case PhaseMenu:
		if inpututil.IsKeyJustPressed(ebiten.Key1) {
			degats := 5 + g.personnage.BonusAttaqueEquipement()
			if structures.EstCoupCritique() {
				degats *= 2
				g.ajouterMessage("*** COUP CRITIQUE ! ***")
			}
			g.combatMonstre.PointsDeVieActuels -= degats
			if g.combatMonstre.PointsDeVieActuels < 0 {
				g.combatMonstre.PointsDeVieActuels = 0
			}
			g.flashMonstre = 10
			g.ajouterMessage(fmt.Sprintf("Tu infliges %d dégâts.", degats))
			g.apresActionJoueur()
		} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
			if len(g.personnage.Sorts) == 0 {
				g.ajouterMessage("Aucun sort connu.")
			} else {
				g.combatPhase = PhaseSorts
			}
		} else if inpututil.IsKeyJustPressed(ebiten.Key3) {
			if len(g.personnage.Inventaire) == 0 {
				g.ajouterMessage("Inventaire vide.")
			} else {
				g.combatPhase = PhaseInventaire
			}
		}
	case PhaseSorts:
		if inpututil.IsKeyJustPressed(ebiten.Key0) {
			g.combatPhase = PhaseMenu
			break
		}
		for i, sortNom := range g.personnage.Sorts {
			if i >= len(touchesNombres) {
				break
			}
			if inpututil.IsKeyJustPressed(touchesNombres[i]) {
				if g.personnage.LancerSort(sortNom, &g.combatMonstre) {
					g.flashMonstre = 10
					g.ajouterMessage("✦ " + sortNom + " !")
					g.apresActionJoueur()
				} else {
					g.ajouterMessage("Mana insuffisant.")
				}
				break
			}
		}
	case PhaseInventaire:
		if inpututil.IsKeyJustPressed(ebiten.Key0) {
			g.combatPhase = PhaseMenu
			break
		}
		for i, item := range g.personnage.Inventaire {
			if i >= len(touchesNombres) {
				break
			}
			if inpututil.IsKeyJustPressed(touchesNombres[i]) {
				g.utiliserObjetCombat(item)
				g.apresActionJoueur()
				break
			}
		}
	case PhaseTourMonstre:
		g.combatMinuteur--
		if g.combatMinuteur <= 0 {
			structures.AttaqueMonstre(&g.combatMonstre, g.personnage, g.combatTourMonstre)
			g.combatTourMonstre++
			g.flashJoueur = 10
			g.ajouterMessage("☠ " + g.combatMonstre.Nom + " attaque !")
			if g.personnage.IsDead() {
				g.ajouterMessage("Tu tombes...")
				g.combatPhase = PhaseDefaite
			} else {
				g.combatPhase = PhaseMenu
			}
		}
	case PhaseVictoire, PhaseDefaite:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.etat = EtatExploration
		}
	}
}

func spriteForMonster(m structures.Monster) (string, int, int) {
	switch m.SpriteID {
	case "zombie":
		return "enemy_zombie.png", 16, 16
	case "knight":
		return "enemy_knight.png", 16, 28
	default:
		return "enemy_pumpkin.png", 16, 23
	}
}

func (g *Game) drawCombat(screen *ebiten.Image) {
	screen.Fill(color.RGBA{10, 10, 17, 255})
	drawTiled(screen, imageAsset("floor.png"), 16, 16, 2)
	vector.DrawFilledRect(screen, 0, 0, 640, 480, color.RGBA{5, 5, 12, 125}, false)

	// Ennemi en grand au centre.
	sprite, w, h := spriteForMonster(g.combatMonstre)
	drawFrame(screen, imageAsset(sprite), g.combatFrame, w, h, 465, 95, 7)
	dessinerPanneau(screen, 20, 18, 600, 72)
	ebitenutil.DebugPrintAt(screen, g.combatMonstre.Nom, 35, 30)
	dessinerBarrePV(screen, 35, 52, 260, 18, g.combatMonstre.PointsDeVieActuels, g.combatMonstre.PointsDeVieMaximum)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("PV %d/%d", g.combatMonstre.PointsDeVieActuels, g.combatMonstre.PointsDeVieMaximum), 305, 55)
	if g.flashMonstre > 0 {
		vector.DrawFilledRect(screen, 420, 75, 180, 150, color.RGBA{255, 255, 255, 70}, false)
	}

	// Joueur.
	dessinerPanneau(screen, 20, 102, 300, 125)
	drawFrame(screen, imageAsset("player_doc.png"), g.combatFrame, 16, 23, 220, 120, 3.5)
	ebitenutil.DebugPrintAt(screen, g.personnage.Nom+"  |  Niv."+fmt.Sprint(g.personnage.Niveau), 35, 116)
	dessinerBarrePV(screen, 35, 138, 155, 14, g.personnage.PointsDeVieActuels, g.personnage.PointsDeVieMaximum)
	dessinerBarreMana(screen, 35, 158, 155, 10, g.personnage.Mana, g.personnage.ManaMax)
	dessinerBarreXP(screen, 35, 177, 155, 8, g.personnage.ExperienceActuelle, g.personnage.ExperienceMax)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Or : %d", g.personnage.Argent), 35, 195)

	// Journal.
	dessinerPanneau(screen, 20, 240, 600, 100)
	for i, msg := range g.combatMessages {
		ebitenutil.DebugPrintAt(screen, msg, 35, 252+i*14)
	}

	// Actions.
	dessinerPanneau(screen, 20, 350, 600, 112)
	switch g.combatPhase {
	case PhaseMenu:
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("[1] ⚔ Attaquer (+%d arme)    [2] ✦ Sorts    [3] 🎒 Inventaire", g.personnage.BonusAttaqueEquipement()), 35, 390)
		ebitenutil.DebugPrintAt(screen, "Chaque action passe le tour à l'ennemi.", 35, 414)
	case PhaseSorts:
		y := 365
		for i, n := range g.personnage.Sorts {
			if i >= len(touchesNombres) {
				break
			}
			s := structures.SortsDisponibles[n]
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("[%d] %s  %ddgts / %dmana", i+1, s.Nom, s.Degats, s.CoutMana), 35, y)
			y += 17
		}
		ebitenutil.DebugPrintAt(screen, "[0] Retour", 35, y)
	case PhaseInventaire:
		y := 365
		for i, item := range g.personnage.Inventaire {
			if i >= len(touchesNombres) {
				break
			}
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("[%d] %s", i+1, item), 35, y)
			y += 17
		}
		ebitenutil.DebugPrintAt(screen, "[0] Retour", 35, y)
	case PhaseTourMonstre:
		ebitenutil.DebugPrintAt(screen, "☠ L'ennemi prépare son attaque...", 35, 395)
	case PhaseVictoire:
		ebitenutil.DebugPrintAt(screen, "✦ VICTOIRE !  ENTREE / ESPACE pour continuer", 35, 395)
	case PhaseDefaite:
		ebitenutil.DebugPrintAt(screen, "☠ K.O.  ENTREE / ESPACE pour continuer", 35, 395)
	}
	if g.combatPopupTimer > 0 {
		dessinerPanneau(screen, 190, 215, 260, 38)
		ebitenutil.DebugPrintAt(screen, g.combatPopup, 210, 230)
	}
}
