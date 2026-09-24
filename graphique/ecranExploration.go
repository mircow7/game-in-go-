package main
 
import (
	"fmt"
	"image/color"
	"math/rand"
 
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)
 
// ============================================================
// Écran d'exploration
// ============================================================
 
func (g *Game) positionnerMonstreAleatoirement() {
	g.monstreX = float64(rand.Intn(largeurEcran - tailleSprite))
	g.monstreY = float64(rand.Intn(hauteurEcran - tailleSprite))
	g.monstreActif = true
}
 
func (g *Game) updateExploration() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.etat = EtatPause
		g.pauseIndex = 0
		return
	}
 
	if !g.monstreActif {
		g.framesAvantRespawn--
		if g.framesAvantRespawn <= 0 {
			g.positionnerMonstreAleatoirement()
		}
	}
 
	vitesse := 4.0
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.joueurX -= vitesse
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.joueurX += vitesse
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.joueurY -= vitesse
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.joueurY += vitesse
	}
 
	if g.joueurX < 0 {
		g.joueurX = 0
	}
	if g.joueurY < 0 {
		g.joueurY = 0
	}
	if g.joueurX > largeurEcran-tailleSprite {
		g.joueurX = largeurEcran - tailleSprite
	}
	if g.joueurY > hauteurEcran-tailleSprite {
		g.joueurY = hauteurEcran - tailleSprite
	}
 
	if g.monstreActif && seChevauchent(g.joueurX, g.joueurY, g.monstreX, g.monstreY, tailleSprite) {
		g.demarrerCombat()
	}
}
 
func (g *Game) drawExploration(screen *ebiten.Image) {
	screen.Fill(color.RGBA{22, 22, 32, 255})
	dessinerGrille(screen)
 
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.joueurX, g.joueurY)
	screen.DrawImage(g.spriteJoueur, opts)
 
	if g.monstreActif {
		optsMonstre := &ebiten.DrawImageOptions{}
		optsMonstre.GeoM.Translate(g.monstreX, g.monstreY)
		screen.DrawImage(g.spriteMonstre, optsMonstre)
	}
 
	// bandeau d'informations en bas de l'écran
	dessinerPanneau(screen, 0, hauteurEcran-36, largeurEcran, 36)
	infos := fmt.Sprintf(
		"%s | Niveau %d | PV %d/%d | Or %d | Echap : menu",
		g.personnage.Nom, g.personnage.Niveau,
		g.personnage.PointsDeVieActuels, g.personnage.PointsDeVieMaximum,
		g.personnage.Argent,
	)
	ebitenutil.DebugPrintAt(screen, infos, 10, hauteurEcran-24)
}
 