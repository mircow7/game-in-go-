package structures

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Génère une initiative aléatoire entre 1 et 20
func RandomInitiative() int {
	return rand.Intn(20) + 1
}

// ============================================================
// MISSION 1 : Initiative — détermine qui commence le combat
// ============================================================

func DeterminerPremierJoueur(p *Personne, m *Monster) string {
	if p.Initiative >= m.Initiative {
		return "joueur"
	}
	return "monstre"
}

// ============================================================
// MISSION 2 : Expérience et montée de niveau
// ============================================================

const ExperienceRequiseBase = 100

func (p *Personne) GainExperience(xpGagne int) {
	fmt.Printf("\nTu gagnes %d points d'expérience !\n", xpGagne)
	p.ExperienceActuelle += xpGagne

	// tant que l'XP dépasse le seuil, on monte de niveau (gère plusieurs
	// niveaux d'un coup si beaucoup d'XP est gagné)
	for p.ExperienceActuelle >= p.ExperienceMax {
		exces := p.ExperienceActuelle - p.ExperienceMax

		p.Niveau++
		p.PointsDeVieMaximum += 20
		p.PointsDeVieActuels = p.PointsDeVieMaximum
		p.ManaMax += 5
		p.Mana = p.ManaMax

		// il faut plus d'XP pour le prochain niveau
		p.ExperienceMax += 50
		p.ExperienceActuelle = exces

		fmt.Printf("*** NIVEAU SUPERIEUR ! Tu es maintenant niveau %d ***\n", p.Niveau)
		fmt.Println("Bonus : +20 PV max, +5 Mana max")
	}

	fmt.Printf("Expérience : %d / %d\n", p.ExperienceActuelle, p.ExperienceMax)
}

// ============================================================
// MISSION 3 : Sorts de combat avec dégâts
// ============================================================

type Sort struct {
	Nom      string
	Degats   int
	CoutMana int
}

var SortsDisponibles = map[string]Sort{
	"Coup de poing": {"Coup de poing", 8, 0},
	"Boule de Feu":  {"Boule de Feu", 18, 10},
}

// MISSION 3 + 4 : lance un sort si le joueur le connait et a assez de mana
func (p *Personne) LancerSort(nomSort string, m *Monster) bool {
	if !p.ConnaitSort(nomSort) {
		fmt.Println("Tu ne connais pas ce sort.")
		return false
	}

	sort, existe := SortsDisponibles[nomSort]
	if !existe {
		fmt.Println("Ce sort n'existe pas.")
		return false
	}

	if p.Mana < sort.CoutMana {
		fmt.Println("Pas assez de mana pour lancer ce sort.")
		return false
	}

	p.Mana -= sort.CoutMana
	m.PointsDeVieActuels -= sort.Degats
	if m.PointsDeVieActuels < 0 {
		m.PointsDeVieActuels = 0
	}

	fmt.Printf("Tu lances %s !\n", sort.Nom)
	fmt.Printf("%s inflige %d dégâts à %s\n", sort.Nom, sort.Degats, m.Nom)
	m.AfficherPV()
	fmt.Printf("Mana restant : %d / %d\n", p.Mana, p.ManaMax)

	return true
}

// ============================================================
// MISSION 4 : Potion de mana
// ============================================================

func (p *Personne) TakeManaPot() {
	if !p.RemoveInventory("Potion de mana") {
		fmt.Println("Tu n'as pas de potion de mana.")
		return
	}
	p.Mana += 30
	if p.Mana > p.ManaMax {
		p.Mana = p.ManaMax
	}
	fmt.Println("Tu bois une potion de mana.")
	fmt.Printf("Mana : %d / %d\n", p.Mana, p.ManaMax)
}

// ============================================================
// MISSION 6 : Easter egg "Qui sont-ils ?"
// ============================================================

func AfficherArtistes() {
	fmt.Println("\n=== Qui sont-ils ? ===")
	fmt.Println("Deux artistes sont cachés dans les parties 2 et 3 du sujet.")
	// TODO : remplace cette ligne par les vrais noms une fois trouvés dans le PDF
	fmt.Println("(noms à compléter une fois identifiés dans les slides)")
}
