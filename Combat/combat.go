package structures

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lecteurCombat = bufio.NewReader(os.Stdin)

func lireCombat() string {
	texte, _ := lecteurCombat.ReadString('\n')
	return strings.TrimSpace(texte)
}

func CharacterTurn(p *Personne, m *Monster) bool {
	fmt.Println("\n--- Ton tour ---")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Print("Ton choix : ")

	switch lireCombat() {
	case "1":
		degats := 5
		m.PointsDeVieActuels -= degats
		if m.PointsDeVieActuels < 0 {
			m.PointsDeVieActuels = 0
		}
		fmt.Printf("Tu utilises Attaque basique\n")
		fmt.Printf("%s inflige %d dégâts à %s\n", p.Nom, degats, m.Nom)
		m.AfficherPV()

	case "2":
		AccessInventoryCombat(p)

	default:
		fmt.Println("Choix invalide, tu perds ton tour.")
	}

	return !m.EstMort()
}

func AccessInventoryCombat(p *Personne) {
	p.AccessInventory()
	fmt.Println("\nTape le numéro d'un objet pour l'utiliser, ou 0 pour ne rien faire.")
	fmt.Print("Ton choix : ")

	choix := lireCombat()
	if choix == "0" {
		return
	}

	index, err := strconv.Atoi(choix)
	if err != nil || index < 1 || index > len(p.Inventaire) {
		fmt.Println("Choix invalide.")
		return
	}

	item := p.Inventaire[index-1]
	switch item {
	case "Potion de vie":
		fmt.Println("Vous utilisez Potion de vie")
		p.TakePot()
	case "Potion de poison":
		fmt.Println("Vous utilisez Potion de poison")
		p.PoisonPot()
	default:
		fmt.Println("Cet objet ne s'utilise pas en combat.")
	}
}

func TrainingFight(p *Personne) {
	gobelin := InitGoblin()
	tour := 1

	fmt.Println("\n===== DEBUT DU COMBAT =====")
	fmt.Println("Un", gobelin.Nom, "apparait !")

	for {
		fmt.Printf("\n=== Tour %d ===\n", tour)

		monstreVivant := CharacterTurn(p, &gobelin)
		if !monstreVivant {
			fmt.Println("\n", gobelin.Nom, "est vaincu !")
			break
		}

		gobelin.GoblinPattern(p, tour)
		if p.IsDead() {
			fmt.Println("Le combat s'arrête ici.")
			break
		}

		tour++
	}

	fmt.Println("\n===== FIN DU COMBAT =====")
}
