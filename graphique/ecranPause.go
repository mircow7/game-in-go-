package main
 
import (
	"image/color"
	"os"
 
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
 
	"game-in-go/structures"
)
 
// ============================================================
// Menu pause (Échap pendant l'exploration)
// ============================================================
 
var optionsPause = []string{"Reprendre", "Sauvegarder la partie", "Quitter le jeu"}
 
func (g *Game) ajouterMessagePause(msg string) {
	g.messagePause = msg
	g.messagePauseMinuteur = 90
}
 
func (g *Game) updatePause() {
	if g.messagePauseMinuteur > 0 {
		g.messagePauseMinuteur--
	}
 
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.pauseIndex--
		if g.pauseIndex < 0 {
			g.pauseIndex = len(optionsPause) - 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.pauseIndex++
		if g.pauseIndex >= len(optionsPause) {
			g.pauseIndex = 0
		}
	}
 
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.etat = EtatExploration
		return
	}
 
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch g.pauseIndex {
		case 0: // Reprendre
			g.etat = EtatExploration
		case 1: // Sauvegarder
			if err := structures.Sauvegarder(*g.personnage); err != nil {
				g.ajouterMessagePause("Erreur lors de la sauvegarde")
			} else {
				g.ajouterMessagePause("Partie sauvegardée !")
			}
		case 2: // Quitter
			os.Exit(0)
		}
	}
}
 
func (g *Game) drawPause(screen *ebiten.Image) {
	// assombrit l'exploration visible en fond
	overlay := ebiten.NewImage(largeurEcran, hauteurEcran)
	overlay.Fill(color.RGBA{0, 0, 0, 150})
	screen.DrawImage(overlay, nil)
 
	largeurPanneau := float32(280)
	hauteurPanneau := float32(200)
	x := float32(largeurEcran)/2 - largeurPanneau/2
	y := float32(hauteurEcran)/2 - hauteurPanneau/2
 
	dessinerPanneau(screen, x, y, largeurPanneau, hauteurPanneau)
	ebitenutil.DebugPrintAt(screen, "PAUSE", int(x)+20, int(y)+20)
 
	for i, option := range optionsPause {
		ligne := "   " + option
		if i == g.pauseIndex {
			ligne = "-> " + option
		}
		ebitenutil.DebugPrintAt(screen, ligne, int(x)+20, int(y)+60+i*24)
	}
 
	if g.messagePauseMinuteur > 0 {
		ebitenutil.DebugPrintAt(screen, g.messagePause, int(x)+20, int(y)+160)
	}
}
 