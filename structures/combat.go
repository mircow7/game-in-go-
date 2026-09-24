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
 
// TACHE 21 + MISSION 3 + BONUS (critique) : tour de jeu du personnage
// Retourne true si le combat continue, false si le monstre est mort
func CharacterTurn(p *Personne, m *Monster) bool {
	fmt.Println("\n--- Ton tour ---")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Sorts")
	fmt.Println("3. Inventaire")
	fmt.Print("Ton choix : ")
 
	switch lireCombat() {
	case "1":
		degats := 5
		if EstCoupCritique() {
			degats *= 2
			fmt.Println("*** COUP CRITIQUE ! ***")
		}
		m.PointsDeVieActuels -= degats
		if m.PointsDeVieActuels < 0 {
			m.PointsDeVieActuels = 0
		}
		fmt.Println("Tu utilises Attaque basique")
		fmt.Printf("%s inflige %d dégâts à %s\n", p.Nom, degats, m.Nom)
		m.AfficherPV()
 
	case "2":
		menuSortsCombat(p, m)
 
	case "3":
		AccessInventoryCombat(p, m)
 
	default:
		fmt.Println("Choix invalide, tu perds ton tour.")
	}
 
	return !m.EstMort()
}
 
// MISSION 3 + 4 : menu des sorts pendant le combat
func menuSortsCombat(p *Personne, m *Monster) {
	if len(p.Sorts) == 0 {
		fmt.Println("Tu ne connais aucun sort.")
		return
	}
 
	fmt.Println("\nTes sorts :")
	for i, sortNom := range p.Sorts {
		sort := SortsDisponibles[sortNom]
		fmt.Printf("%d. %s (%d dégâts, %d mana)\n", i+1, sort.Nom, sort.Degats, sort.CoutMana)
	}
	fmt.Println("0. Annuler")
	fmt.Print("Ton choix : ")
 
	choix := lireCombat()
	if choix == "0" {
		return
	}
 
	index, err := strconv.Atoi(choix)
	if err != nil || index < 1 || index > len(p.Sorts) {
		fmt.Println("Choix invalide.")
		return
	}
 
	p.LancerSort(p.Sorts[index-1], m)
}
 
// TACHE 21 suite + BONUS (AK-47) : utiliser un objet de l'inventaire pendant le combat
func AccessInventoryCombat(p *Personne, m *Monster) {
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
	case "Potion de mana":
		fmt.Println("Vous utilisez Potion de mana")
		p.TakeManaPot()
	case "AK-47":
		UtiliserAK47(p, m)
	default:
		fmt.Println("Cet objet ne s'utilise pas en combat.")
	}
}
 
// BONUS : arme trouvée sur un Troll, clairement pas d'époque, dégâts énormes, usage unique
func UtiliserAK47(p *Personne, m *Monster) {
	if !p.RemoveInventory("AK-47") {
		return
	}
	degats := 35
	m.PointsDeVieActuels -= degats
	if m.PointsDeVieActuels < 0 {
		m.PointsDeVieActuels = 0
	}
	fmt.Println("Tu sors un AK-47... personne ne sait comment c'est arrivé ici.")
	fmt.Printf("Rafale ! %s inflige %d dégâts à %s\n", p.Nom, degats, m.Nom)
	m.AfficherPV()
}
 
// TACHE 22 + MISSION 1 + MISSION 2 + BONUS (monstre aléatoire, critique) : combat complet
func TrainingFight(p *Personne) {
	DeroulerCombat(p, ChoisirMonsterAleatoire())
}
 
// BONUS : logique de combat réutilisable (Entrainement et Donjon l'utilisent tous les deux)
// Retourne true si le joueur a gagné, false s'il est mort en cours de route.
func DeroulerCombat(p *Personne, monstre Monster) bool {
	tour := 1
 
	fmt.Println("\n===== DEBUT DU COMBAT =====")
	fmt.Println(ObtenirArtMonstre(monstre.Nom))
	fmt.Println("Un", monstre.Nom, "apparait !")
 
	// MISSION 1 : qui commence le combat ?
	premier := DeterminerPremierJoueur(p, &monstre)
	if premier == "joueur" {
		fmt.Println("Ton initiative est plus haute, tu commences !")
	} else {
		fmt.Println(monstre.Nom, "est plus rapide, il attaque en premier !")
	}
 
	for {
		fmt.Printf("\n=== Tour %d ===\n", tour)
 
		if premier == "joueur" {
			if !CharacterTurn(p, &monstre) {
				break
			}
			AttaqueMonstre(&monstre, p, tour)
			if p.IsDead() {
				AfficherDefaite()
				fmt.Println("Le combat s'arrête ici.")
				return false
			}
		} else {
			AttaqueMonstre(&monstre, p, tour)
			if p.IsDead() {
				AfficherDefaite()
				fmt.Println("Le combat s'arrête ici.")
				return false
			}
			if !CharacterTurn(p, &monstre) {
				break
			}
		}
 
		tour++
	}
 
	fmt.Println("\n", monstre.Nom, "est vaincu !")
	AfficherVictoire()
 
	// MISSION 2 : récompense d'expérience
	p.GainExperience(monstre.ExperienceOffert)
 
	// BONUS : butin aléatoire
	if objet, obtenu := TirerButin(monstre); obtenu {
		if p.InventairePlein() {
			fmt.Println("Le monstre a lâché", objet, "mais ton inventaire est plein !")
		} else {
			fmt.Println("Le monstre a lâché un objet :", objet)
			p.AddInventory(objet)
		}
	}
 
	fmt.Println("\n===== FIN DU COMBAT =====")
	return true
}
 
// BONUS : coup critique possible côté monstre, en plus du pattern habituel
func AttaqueMonstre(m *Monster, p *Personne, tour int) {
	if EstCoupCritique() {
		fmt.Println("***", m.Nom, "place un coup critique ! ***")
		degatsNormaux := m.PointsAttaque
		m.PointsAttaque *= 2
		m.GoblinPattern(p, tour)
		m.PointsAttaque = degatsNormaux
		return
	}
	m.GoblinPattern(p, tour)
}
 