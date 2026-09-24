package main
 
import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os">
	"strings"
 
	"github.com/hajimehoshi/ebiten/v2"
 
	"game-in-go/structures"
)
 
// ============================================================
// Constantes générales de la fenêtre
// ============================================================
 
const (
	largeurEcran = 640
	hauteurEcran = 480
	tailleSprite = 32
)
 
// EtatJeu représente le grand écran affiché à l'instant T
type EtatJeu int
 
const (
	EtatTitre EtatJeu = iota
	EtatExploration
	EtatCombat
	EtatPause
)
 
// ============================================================
// Structure principale du jeu
// ============================================================
 
// Game est l'état complet du jeu. Ebiten appelle Update() puis Draw()
// automatiquement, environ 60 fois par seconde.
type Game struct {
	etat EtatJeu
 
	personnage *structures.Personne
 
	// --- exploration ---
	joueurX, joueurY   float64
	spriteJoueur       *ebiten.Image
	spriteMonstre      *ebiten.Image
	monstreX, monstreY float64
	monstreActif       bool
	framesAvantRespawn int
 
	// --- combat ---
	combatMonstre     structures.Monster
	combatPhase       PhaseCombat
	combatTourMonstre int
	combatMessages    []string
	combatMinuteur    int
	flashJoueur       int
	flashMonstre      int
 
	// --- pause ---
	pauseIndex           int
	messagePause         string
	messagePauseMinuteur int
}
 
// ============================================================
// Boucle principale Ebiten
// ============================================================
 
func (g *Game) Update() error {
	switch g.etat {
	case EtatTitre:
		g.updateTitre()
	case EtatExploration:
		g.updateExploration()
	case EtatCombat:
		g.updateCombat()
	case EtatPause:
		g.updatePause()
	}
	return nil
}
 
func (g *Game) Draw(screen *ebiten.Image) {
	switch g.etat {
	case EtatTitre:
		g.drawTitre(screen)
	case EtatExploration:
		g.drawExploration(screen)
	case EtatCombat:
		g.drawCombat(screen)
	case EtatPause:
		g.drawExploration(screen) // on voit l'exploration en fond, assombrie
		g.drawPause(screen)
	}
}
 
func (g *Game) Layout(largeurExterne, hauteurExterne int) (int, int) {
	return largeurEcran, hauteurEcran
}
 
// ============================================================
// Point d'entrée
// ============================================================
 
func main() {
	fmt.Println("=== PROJET RED - Version graphique ===")
 
	var personnage structures.Personne
 
	if structures.SauvegardeExiste() {
		fmt.Println("Une sauvegarde a été trouvée.")
		fmt.Println("1. Charger la partie sauvegardée")
		fmt.Println("2. Commencer une nouvelle partie")
		fmt.Print("Ton choix : ")
 
		lecteur := bufio.NewReader(os.Stdin)
		reponse, _ := lecteur.ReadString('\n')
		reponse = strings.TrimSpace(reponse)
 
		if reponse == "1" {
			p, ok := structures.Charger()
			if ok {
				personnage = p
				fmt.Println("Partie chargée !")
			} else {
				fmt.Println("Impossible de charger la sauvegarde, nouvelle partie.")
				personnage = structures.CharacterCreation()
			}
		} else {
			personnage = structures.CharacterCreation()
		}
	} else {
		personnage = structures.CharacterCreation()
	}
 
	fmt.Println("\nUne fenêtre de jeu va s'ouvrir. Bonne aventure !")
 
	ebiten.SetWindowSize(largeurEcran, hauteurEcran)
	ebiten.SetWindowTitle("Projet RED")
 
	jeu := &Game{
		etat:          EtatTitre,
		personnage:    &personnage,
		joueurX:       100,
		joueurY:       100,
		spriteJoueur:  nouvelleImageCouleur(tailleSprite, color.RGBA{80, 140, 255, 255}),
		spriteMonstre: nouvelleImageCouleur(tailleSprite, color.RGBA{200, 60, 60, 255}),
	}
	jeu.positionnerMonstreAleatoirement()
 
	if err := ebiten.RunGame(jeu); err != nil {
		log.Fatal(err)
	}
}
 