package main
 
import (
	"fmt"
	"image/color"
 
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
 
	"game-in-go/structures"
)
 
// ============================================================
// Écran de combat
// ============================================================
 
// PhaseCombat détaille où on en est à l'intérieur d'un combat
type PhaseCombat int
 
const (
	PhaseMenu PhaseCombat = iota
	PhaseSorts
	PhaseInventaire
	PhaseTourMonstre
	PhaseVictoire
	PhaseDefaite
)
 
// touchesNombres permet de choisir un élément de liste (sort, objet) au clavier (1 à 9)
var touchesNombres = []ebiten.Key{
	ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.Key5,
	ebiten.Key6, ebiten.Key7, ebiten.Key8, ebiten.Key9,
}
 
func (g *Game) ajouterMessage(msg string) {
	g.combatMessages = append(g.combatMessages, msg)
	if len(g.combatMessages) > 6 {
		g.combatMessages = g.combatMessages[len(g.combatMessages)-6:]
	}
}
 
func (g *Game) demarrerCombat() {
	g.combatMonstre = structures.ChoisirMonsterAleatoire()
	g.etat = EtatCombat
	g.combatMessages = nil
	g.combatTourMonstre = 1
	g.combatPhase = PhaseMenu
	g.flashJoueur = 0
	g.flashMonstre = 0
 
	g.ajouterMessage(fmt.Sprintf("Un %s apparait !", g.combatMonstre.Nom))
 
	premier := structures.DeterminerPremierJoueur(g.personnage, &g.combatMonstre)
	if premier != "joueur" {
		g.ajouterMessage(g.combatMonstre.Nom + " est plus rapide !")
		g.combatPhase = PhaseTourMonstre
		g.combatMinuteur = 40
	} else {
		g.ajouterMessage("Tu es plus rapide, tu commences !")
	}
}
 
// apresActionJoueur vérifie si le monstre est mort, sinon passe la main au monstre
func (g *Game) apresActionJoueur() {
	if g.combatMonstre.EstMort() {
		g.ajouterMessage(g.combatMonstre.Nom + " est vaincu !")
		g.personnage.GainExperience(g.combatMonstre.ExperienceOffert)
 
		if objet, obtenu := structures.TirerButin(g.combatMonstre); obtenu {
			if g.personnage.InventairePlein() {
				g.ajouterMessage("Butin trouvé (" + objet + ") mais inventaire plein !")
			} else {
				g.personnage.AddInventory(objet)
				g.ajouterMessage("Butin obtenu : " + objet)
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
		g.ajouterMessage("Tu utilises une Potion de vie")
	case "Potion de poison":
		g.personnage.PoisonPot()
		g.ajouterMessage("Tu utilises une Potion de poison")
	case "Potion de mana":
		g.personnage.TakeManaPot()
		g.ajouterMessage("Tu utilises une Potion de mana")
	case "AK-47":
		structures.UtiliserAK47(g.personnage, &g.combatMonstre)
		g.flashMonstre = 10
		g.ajouterMessage("Tu sors un AK-47 !")
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
 
	switch g.combatPhase {
 
	case PhaseMenu:
		if inpututil.IsKeyJustPressed(ebiten.Key1) {
			degats := 5
			if structures.EstCoupCritique() {
				degats *= 2
				g.ajouterMessage("*** COUP CRITIQUE ! ***")
			}
			g.combatMonstre.PointsDeVieActuels -= degats
			if g.combatMonstre.PointsDeVieActuels < 0 {
				g.combatMonstre.PointsDeVieActuels = 0
			}
			g.flashMonstre = 10
			g.ajouterMessage(fmt.Sprintf("Tu infliges %d dégâts à %s", degats, g.combatMonstre.Nom))
			g.apresActionJoueur()
 
		} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
			if len(g.personnage.Sorts) == 0 {
				g.ajouterMessage("Tu ne connais aucun sort.")
			} else {
				g.combatPhase = PhaseSorts
			}
 
		} else if inpututil.IsKeyJustPressed(ebiten.Key3) {
			if len(g.personnage.Inventaire) == 0 {
				g.ajouterMessage("Ton inventaire est vide.")
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
				succes := g.personnage.LancerSort(sortNom, &g.combatMonstre)
				if succes {
					g.flashMonstre = 10
					g.ajouterMessage("Tu lances " + sortNom)
					g.apresActionJoueur()
				} else {
					g.ajouterMessage("Impossible (mana insuffisant ?)")
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
			g.ajouterMessage(fmt.Sprintf("%s attaque !", g.combatMonstre.Nom))
 
			if g.personnage.IsDead() {
				g.ajouterMessage("Tu tombes... mais tu te relèves avec la moitié de tes PV.")
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
 
func (g *Game) drawCombat(screen *ebiten.Image) {
	screen.Fill(color.RGBA{15, 15, 22, 255})
 
	// --- panneau monstre ---
	dessinerPanneau(screen, 20, 20, largeurEcran-40, 70)
	ebitenutil.DebugPrintAt(screen, g.combatMonstre.Nom, 32, 32)
	dessinerBarrePV(screen, 32, 55, 260, 18, g.combatMonstre.PointsDeVieActuels, g.combatMonstre.PointsDeVieMaximum)
	if g.flashMonstre > 0 {
		vector.DrawFilledRect(screen, 20, 20, largeurEcran-40, 70, color.RGBA{255, 255, 255, 60}, false)
	}
 
	// --- journal de combat ---
	dessinerPanneau(screen, 20, 100, largeurEcran-40, 130)
	for i, msg := range g.combatMessages {
		ebitenutil.DebugPrintAt(screen, msg, 32, 112+i*17)
	}
 
	// --- panneau joueur ---
	dessinerPanneau(screen, 20, 245, largeurEcran-40, 90)
	ebitenutil.DebugPrintAt(screen, g.personnage.Nom, 32, 257)
	dessinerBarrePV(screen, 32, 280, 260, 18, g.personnage.PointsDeVieActuels, g.personnage.PointsDeVieMaximum)
	dessinerBarreMana(screen, 32, 304, 260, 14, g.personnage.Mana, g.personnage.ManaMax)
	if g.flashJoueur > 0 {
		vector.DrawFilledRect(screen, 20, 245, largeurEcran-40, 90, color.RGBA{255, 60, 60, 70}, false)
	}
 
	// --- panneau d'action (menu / sous-menus) ---
	dessinerPanneau(screen, 20, 345, largeurEcran-40, 115)
 
	switch g.combatPhase {
	case PhaseMenu:
		ebitenutil.DebugPrintAt(screen, "1: Attaquer   2: Sorts   3: Inventaire", 32, 400)
 
	case PhaseSorts:
		y := 358
		for i, sortNom := range g.personnage.Sorts {
			sort := structures.SortsDisponibles[sortNom]
			ebitenutil.DebugPrintAt(screen,
				fmt.Sprintf("%d: %s (%d dgts, %d mana)", i+1, sort.Nom, sort.Degats, sort.CoutMana),
				32, y)
			y += 16
		}
		ebitenutil.DebugPrintAt(screen, "0: Retour", 32, y)
 
	case PhaseInventaire:
		y := 358
		for i, item := range g.personnage.Inventaire {
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d: %s", i+1, item), 32, y)
			y += 16
		}
		ebitenutil.DebugPrintAt(screen, "0: Retour", 32, y)
 
	case PhaseTourMonstre:
		ebitenutil.DebugPrintAt(screen, "...", 32, 400)
 
	case PhaseVictoire:
		ebitenutil.DebugPrintAt(screen, "VICTOIRE ! Appuie sur ENTREE pour continuer.", 32, 400)
 
	case PhaseDefaite:
		ebitenutil.DebugPrintAt(screen, "K.O... Appuie sur ENTREE pour continuer.", 32, 400)
	}
}
 