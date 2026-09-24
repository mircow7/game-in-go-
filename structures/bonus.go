package structures

import (
	"fmt"
	"math/rand"
)

// ============================================================
// BONUS : coups critiques (10% de chance, dégâts x2)
// ============================================================

func EstCoupCritique() bool {
	return rand.Intn(100) < 10
}

// ============================================================
// BONUS : butin aléatoire à la mort d'un monstre
// ============================================================

// TauxDeDrop : chance en % qu'un monstre lâche un objet en mourant
const TauxDeDrop = 40

// TirerButin renvoie un objet au hasard dans le butin du monstre, et true
// si le tirage est réussi. Renvoie ("", false) si rien n'est obtenu ou si
// le monstre n'a pas de butin défini.
func TirerButin(m Monster) (string, bool) {
	if len(m.Butin) == 0 {
		return "", false
	}
	if rand.Intn(100) >= TauxDeDrop {
		return "", false
	}
	objet := m.Butin[rand.Intn(len(m.Butin))]
	return objet, true
}

// ============================================================
// BONUS : revente d'objets au marchand (à moitié prix)
// ============================================================

// Doit correspondre aux prix du marchand dans main.go
var PrixObjets = map[string]int{
	"Potion de vie":                3,
	"Potion de poison":             6,
	"Potion de mana":               8,
	"Livre de Sort : Boule de Feu": 25,
	"Fourrure de Loup":             4,
	"Peau de Troll":                7,
	"Cuir de Sanglier":             3,
	"Plume de Corbeau":             1,
}

// Vend un objet de l'inventaire à la moitié de son prix d'achat.
// Retourne false si l'objet n'est pas dans l'inventaire ou n'a pas de prix connu.
func (p *Personne) VendreObjet(nomObjet string) bool {
	prix, existe := PrixObjets[nomObjet]
	if !existe {
		fmt.Println("Cet objet ne peut pas être revendu.")
		return false
	}

	if !p.RemoveInventory(nomObjet) {
		fmt.Println("Tu n'as pas cet objet dans ton inventaire.")
		return false
	}

	gain := prix / 2
	if gain < 1 {
		gain = 1
	}
	p.Argent += gain

	fmt.Printf("Tu vends %s pour %d pièces d'or.\n", nomObjet, gain)
	return true
}
