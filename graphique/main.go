package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"
	"strings"

	"game-in-go/structures"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	largeurEcran = 640
	hauteurEcran = 480
	tailleSprite = 32
)

type EtatJeu int

const (
	EtatTitre EtatJeu = iota
	EtatAccueil
	EtatExploration
	EtatCombat
	EtatPause
	EtatInventaire
	EtatBoutique
	EtatEquipement
	EtatCraft
)

type Game struct {
	etat       EtatJeu
	personnage *structures.Personne

	// Exploration
	joueurX, joueurY   float64
	monstreX, monstreY float64
	monstreActif       bool
	monstreExploration structures.Monster
	framesAvantRespawn int
	frameAnim          int
	animTick           int

	// Coffre / or
	coffreX, coffreY float64
	coffreActif      bool
	coinX, coinY     float64
	coinActif        bool
	coinTick         int

	// Combat
	combatMonstre     structures.Monster
	combatPhase       PhaseCombat
	combatTourMonstre int
	combatMessages    []string
	combatMinuteur    int
	flashJoueur       int
	flashMonstre      int
	combatAnimTick    int
	combatFrame       int
	combatPopup       string
	combatPopupTimer  int

	// Accueil / hub
	accueilIndex        int
	accueilMessage      string
	accueilMessageTimer int

	// Inventaire
	inventaireIndex        int
	inventaireMessage      string
	inventaireMessageTimer int

	// Boutique
	boutiquePage         int
	boutiqueIndex        int
	boutiqueMessage      string
	boutiqueMessageTimer int

	// Equipement
	equipementIndex        int
	equipementMessage      string
	equipementMessageTimer int

	// Crafting
	craftIndex        int
	craftMessage      string
	craftMessageTimer int

	// Pause
	pauseIndex           int
	messagePause         string
	messagePauseMinuteur int
}

func (g *Game) Update() error {
	switch g.etat {
	case EtatTitre:
		g.updateTitre()
	case EtatAccueil:
		g.updateAccueil()
	case EtatExploration:
		g.updateExploration()
	case EtatCombat:
		g.updateCombat()
	case EtatPause:
		g.updatePause()
	case EtatInventaire:
		g.updateInventaire()
	case EtatBoutique:
		g.updateBoutique()
	case EtatEquipement:
		g.updateEquipement()
	case EtatCraft:
		g.updateCraft()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.etat {
	case EtatTitre:
		g.drawTitre(screen)
	case EtatAccueil:
		g.drawAccueil(screen)
	case EtatExploration:
		g.drawExploration(screen)
	case EtatCombat:
		g.drawCombat(screen)
	case EtatPause:
		g.drawExploration(screen)
		g.drawPause(screen)
	case EtatInventaire:
		g.drawInventaire(screen)
	case EtatBoutique:
		g.drawBoutique(screen)
	case EtatEquipement:
		g.drawEquipement(screen)
	case EtatCraft:
		g.drawCraft(screen)
	}
}

func (g *Game) Layout(_, _ int) (int, int) { return largeurEcran, hauteurEcran }

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
			} else {
				personnage = structures.CharacterCreation()
			}
		} else {
			personnage = structures.CharacterCreation()
		}
	} else {
		personnage = structures.CharacterCreation()
	}

	chargerAssets()

	ebiten.SetWindowSize(largeurEcran, hauteurEcran)
	ebiten.SetWindowTitle("Projet RED — Dungeon Edition")

	jeu := &Game{
		etat: EtatTitre, personnage: &personnage,
		joueurX: 90, joueurY: 95,
		monstreActif: true,
		coffreActif:  true, coffreX: 520, coffreY: 150,
		coinActif: true, coinX: 310, coinY: 250,
	}
	jeu.positionnerMonstreAleatoirement()

	if err := ebiten.RunGame(jeu); err != nil {
		log.Fatal(err)
	}
}

var _ = color.RGBA{}
