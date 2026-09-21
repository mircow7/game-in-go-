package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
 
	"game-in-go/structures"
)
 
var lecteur = bufio.NewReader(os.Stdin)
 
// Lit une ligne tapée par le joueur et enlève le retour à la ligne
func lire() string {
	texte, _ := lecteur.ReadString('\n')
	return strings.TrimSpace(texte)
}
 
func main() {
	var c1 structures.Personne
 
	// BONUS : proposer de charger une sauvegarde existante
	if structures.SauvegardeExiste() {
		fmt.Println("Une sauvegarde a été trouvée.")
		fmt.Println("1. Charger la partie sauvegardée")
		fmt.Println("2. Commencer une nouvelle partie")
		fmt.Print("Ton choix : ")
 
		if lire() == "1" {
			perso, ok := structures.Charger()
			if ok {
				c1 = perso
				fmt.Println("Partie chargée !")
			} else {
				fmt.Println("Impossible de charger la sauvegarde, nouvelle partie.")
				c1 = structures.CharacterCreation()
			}
		} else {
			c1 = structures.CharacterCreation()
		}
	} else {
		// TACHE 11 : le joueur crée lui-même son personnage
		c1 = structures.CharacterCreation()
	}
 
	// TACHE 6 : menu principal
	// TACHE 15 : + Forgeron   TACHE 22 : + Entrainement   BONUS : + Vendre, Sauvegarder
	for {
		fmt.Println("\n===== MENU PRINCIPAL =====")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder au contenu de l'inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Forgeron")
		fmt.Println("5. Entrainement")
		fmt.Println("6. Donjon")
		fmt.Println("7. Qui sont-ils ?")
		fmt.Println("8. Vendre un objet")
		fmt.Println("9. Sauvegarder la partie")
		fmt.Println("10. Quitter")
		fmt.Print("Ton choix : ")
 
		switch lire() {
		case "1":
			c1.DisplayInfo()
		case "2":
			menuInventaire(&c1)
		case "3":
			menuMarchand(&c1)
		case "4":
			menuForgeron(&c1)
		case "5":
			structures.TrainingFight(&c1)
		case "6":
			structures.LancerDonjon(&c1)
		case "7":
			structures.AfficherArtistes()
		case "8":
			menuVente(&c1)
		case "9":
			if err := structures.Sauvegarder(c1); err != nil {
				fmt.Println("Erreur lors de la sauvegarde :", err)
			} else {
				fmt.Println("Partie sauvegardée !")
			}
		case "10":
			fmt.Println("À bientôt !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
 
// BONUS : revendre un objet de l'inventaire
func menuVente(p *structures.Personne) {
	for {
		p.AccessInventory()
		fmt.Println("\nTape le numéro d'un objet à vendre, ou 0 pour revenir.")
		fmt.Print("Ton choix : ")
 
		choix := lire()
		if choix == "0" {
			return
		}
 
		index, err := strconv.Atoi(choix)
		if err != nil || index < 1 || index > len(p.Inventaire) {
			fmt.Println("Choix invalide.")
			continue
		}
 
		p.VendreObjet(p.Inventaire[index-1])
	}
}
 
// TACHE 4 + 5 : afficher l'inventaire et utiliser un objet
func menuInventaire(p *structures.Personne) {
	for {
		p.AccessInventory()
		fmt.Println("\nTape le numéro d'un objet pour l'utiliser, ou 0 pour revenir.")
		fmt.Print("Ton choix : ")
		choix := lire()
 
		if choix == "0" {
			return
		}
 
		index, err := strconv.Atoi(choix)
		if err != nil || index < 1 || index > len(p.Inventaire) {
			fmt.Println("Choix invalide.")
			continue
		}
 
		item := p.Inventaire[index-1]
		switch item {
		case "Potion de vie":
			p.TakePot()
		case "Potion de poison":
			p.PoisonPot()
		case "Livre de Sort : Boule de Feu":
			if p.RemoveInventory(item) {
				p.SpellBook("Boule de Feu")
			}
		case "Chapeau de l'aventurier":
			p.Equiper(item, "tete")
		case "Tunique de l'aventurier":
			p.Equiper(item, "torse")
		case "Bottes de l'aventurier":
			p.Equiper(item, "pieds")
		default:
			fmt.Println("Cet objet ne s'utilise pas encore.")
		}
	}
}
 
// TACHE 7 + 9 + 14 : interface du marchand avec prix en pièces d'or
type ArticleMarchand struct {
	Nom  string
	Prix int
}
 
var articlesMarchand = []ArticleMarchand{
	{"Potion de vie", 3},
	{"Potion de poison", 6},
	{"Potion de mana", 8},
	{"Livre de Sort : Boule de Feu", 25},
	{"Fourrure de Loup", 4},
	{"Peau de Troll", 7},
	{"Cuir de Sanglier", 3},
	{"Plume de Corbeau", 1},
	{"Augmentation d'inventaire", 30},
}
 
func menuMarchand(p *structures.Personne) {
	for {
		fmt.Println("\n===== MARCHAND =====")
		for i, a := range articlesMarchand {
			fmt.Printf("%d. %s (%d pièces d'or)\n", i+1, a.Nom, a.Prix)
		}
		fmt.Println("0. Retour")
		fmt.Printf("Ton or : %d\n", p.Argent)
		fmt.Print("Ton choix : ")
 
		choix := lire()
		if choix == "0" {
			return
		}
 
		index, err := strconv.Atoi(choix)
		if err != nil || index < 1 || index > len(articlesMarchand) {
			fmt.Println("Choix invalide.")
			continue
		}
 
		article := articlesMarchand[index-1]
 
		if p.Argent < article.Prix {
			fmt.Println("Tu n'as pas assez d'or pour acheter :", article.Nom)
			continue
		}
 
		// TACHE 18 : cas particulier, ce n'est pas un objet d'inventaire
		if article.Nom == "Augmentation d'inventaire" {
			p.Argent -= article.Prix
			p.UpgradeInventorySlot()
			continue
		}
 
		if p.InventairePlein() {
			fmt.Println("Ton inventaire est plein.")
			continue
		}
 
		p.Argent -= article.Prix
		p.AddInventory(article.Nom)
	}
}
 
// TACHE 15 : interface du forgeron
func menuForgeron(p *structures.Personne) {
	objets := []string{
		"Chapeau de l'aventurier",
		"Tunique de l'aventurier",
		"Bottes de l'aventurier",
	}
 
	for {
		fmt.Println("\n===== FORGERON =====")
		for i, nom := range objets {
			recette := structures.Recettes[nom]
			fmt.Printf("%d. %s (%d pièces d'or)\n", i+1, nom, recette.Prix)
		}
		fmt.Println("0. Retour")
		fmt.Print("Ton choix : ")
 
		choix := lire()
		if choix == "0" {
			return
		}
 
		index, err := strconv.Atoi(choix)
		if err != nil || index < 1 || index > len(objets) {
			fmt.Println("Choix invalide.")
			continue
		}
 
		p.Fabriquer(objets[index-1])
	}
}
 