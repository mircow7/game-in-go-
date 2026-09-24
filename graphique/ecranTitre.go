package main
 
import (
	"fmt"
	"image/color"
 
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)
 
// ============================================================
// Écran titre
// ============================================================
 
func (g *Game) updateTitre() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.etat = EtatExploration
	}
}
 
func (g *Game) drawTitre(screen *ebiten.Image) {
	screen.Fill(color.RGBA{18, 18, 28, 255})
 
	dessinerPanneau(screen, 90, 130, largeurEcran-180, 220)
 
	ebitenutil.DebugPrintAt(screen, "===================================", 110, 150)
	ebitenutil.DebugPrintAt(screen, "            PROJET RED             ", 110, 170)
	ebitenutil.DebugPrintAt(screen, "===================================", 110, 190)
 
	infosPerso := fmt.Sprintf("%s le %s - Niveau %d", g.personnage.Nom, g.personnage.Classe, g.personnage.Niveau)
	ebitenutil.DebugPrintAt(screen, infosPerso, 110, 230)
 
	ebitenutil.DebugPrintAt(screen, "Appuie sur ENTREE pour commencer l'aventure", 110, 270)
	ebitenutil.DebugPrintAt(screen, "Fleches : se deplacer", 110, 300)
	ebitenutil.DebugPrintAt(screen, "Echap : menu pause (sauvegarder, quitter)", 110, 320)
}
 