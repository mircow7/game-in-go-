package structures

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

// TACHE 12 : limite d'inventaire
const TailleMaxInventaire = 10

// TACHE 16 : équipement du personnage
type Equipment struct {
	Tete  string
	Torse string
	Pieds string
	Arme  string
}

// TACHE 1 : structure du personnage
// TACHE 10 : + Sorts   TACHE 13 : + Argent   TACHE 16 : + Equipement
type Personne struct {
	Nom                string
	Classe             string
	Niveau             int
	PointsDeVieMaximum int
	PointsDeVieActuels int
	Inventaire         []string
	Sorts              []string
	Argent             int
	Equipement         Equipment

	// TACHE 18 : nombre d'améliorations d'inventaire déjà utilisées
	AmeliorationsInventaire int

	// MISSION 1 : initiative pour déterminer qui commence le combat
	Initiative int

	// MISSION 2 : expérience et niveau
	ExperienceActuelle int
	ExperienceMax      int

	// MISSION 4 : mana pour lancer les sorts
	Mana    int
	ManaMax int
}

// TACHE 2 : initialisation du personnage
// TACHE 10 : ajoute le sort de base "Coup de poing"
// TACHE 13 : donne 100 pièces d'or de départ
func InitCharacter(nom string, classe string, niveau int, pvMax int, pvActuels int, inventaire []string) Personne {
	return Personne{
		Nom:                nom,
		Classe:             classe,
		Niveau:             niveau,
		PointsDeVieMaximum: pvMax,
		PointsDeVieActuels: pvActuels,
		Inventaire:         inventaire,
		Sorts:              []string{"Coup de poing"},
		Argent:             100,
	}
}

// TACHE 12 : la taille max de l'inventaire dépend des améliorations achetées
func (p Personne) TailleMaxInventaireActuelle() int {
	return TailleMaxInventaire + p.AmeliorationsInventaire*10
}

// TACHE 3 : affichage des informations
// MISSION 2 + 4 : affiche aussi l'expérience et le mana
func (p Personne) DisplayInfo() {
	fmt.Println("\n=== Informations du personnage ===")
	fmt.Println("Nom    :", p.Nom)
	fmt.Println("Classe :", p.Classe)
	fmt.Println("Niveau :", p.Niveau)
	p.AfficherPV()
	fmt.Printf("Mana : %d / %d\n", p.Mana, p.ManaMax)
	fmt.Printf("Expérience : %d / %d\n", p.ExperienceActuelle, p.ExperienceMax)
	fmt.Printf("Inventaire : %d / %d objets\n", len(p.Inventaire), p.TailleMaxInventaireActuelle())
}

func (p Personne) AfficherPV() {
	fmt.Printf("PV : %d / %d\n", p.PointsDeVieActuels, p.PointsDeVieMaximum)
}

// TACHE 4 : accès à l'inventaire (affichage seulement)
func (p Personne) AccessInventory() {
	fmt.Println("\n=== Inventaire ===")
	if len(p.Inventaire) == 0 {
		fmt.Println("(vide)")
		return
	}
	for i, item := range p.Inventaire {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}

// TACHE 7 : fonctions génériques d'ajout / retrait
// TACHE 12 : utilise désormais la taille max variable
func (p Personne) InventairePlein() bool {
	return len(p.Inventaire) >= p.TailleMaxInventaireActuelle()
}

func (p *Personne) AddInventory(item string) {
	if p.InventairePlein() {
		fmt.Println("Ton sac est plein, impossible d'ajouter :", item)
		return
	}
	p.Inventaire = append(p.Inventaire, item)
	fmt.Println("Ajouté à l'inventaire :", item)
}

// Renvoie true si l'objet a bien été retiré
func (p *Personne) RemoveInventory(item string) bool {
	for i, it := range p.Inventaire {
		if it == item {
			p.Inventaire = append(p.Inventaire[:i], p.Inventaire[i+1:]...)
			return true
		}
	}
	return false
}

// TACHE 5 : potion de vie
func (p *Personne) TakePot() {
	if !p.RemoveInventory("Potion de vie") {
		fmt.Println("Tu n'as pas de potion de vie.")
		return
	}
	p.PointsDeVieActuels += 50
	if p.PointsDeVieActuels > p.PointsDeVieMaximum {
		p.PointsDeVieActuels = p.PointsDeVieMaximum
	}
	fmt.Println("Tu bois une potion de vie.")
	p.AfficherPV()
}

// TACHE 9 : potion de poison
func (p *Personne) PoisonPot() {
	if !p.RemoveInventory("Potion de poison") {
		fmt.Println("Tu n'as pas de potion de poison.")
		return
	}
	fmt.Println("Tu bois la potion... c'était du poison !")
	for tour := 1; tour <= 3; tour++ {
		time.Sleep(1 * time.Second)
		p.PointsDeVieActuels -= 10
		if p.PointsDeVieActuels < 0 {
			p.PointsDeVieActuels = 0
		}
		p.AfficherPV()
		if p.IsDead() {
			return
		}
	}
}

// TACHE 8 : mort et résurrection
func (p *Personne) IsDead() bool {
	if p.PointsDeVieActuels > 0 {
		return false
	}
	fmt.Println("\n*** WASTED ***")
	fmt.Println("Tu es ressuscité avec la moitié de tes points de vie.")
	p.PointsDeVieActuels = p.PointsDeVieMaximum / 2
	p.AfficherPV()
	return true
}

// TACHE 10 : apprentissage de sorts
func (p Personne) ConnaitSort(sort string) bool {
	for _, s := range p.Sorts {
		if s == sort {
			return true
		}
	}
	return false
}

func (p *Personne) SpellBook(sort string) {
	if p.ConnaitSort(sort) {
		fmt.Println("Tu connais déjà le sort :", sort)
		return
	}
	p.Sorts = append(p.Sorts, sort)
	fmt.Println("Nouveau sort appris :", sort)
}

// TACHE 17 : équiper un objet fabriqué
// section doit valoir "tete", "torse" ou "pieds"
func (p *Personne) Equiper(nomObjet string, section string) {
	if !p.RemoveInventory(nomObjet) {
		fmt.Println("Tu n'as pas cet objet dans ton inventaire.")
		return
	}

	var ancien string
	var bonusPV int

	switch section {
	case "tete":
		ancien = p.Equipement.Tete
		p.Equipement.Tete = nomObjet
		bonusPV = 10
	case "torse":
		ancien = p.Equipement.Torse
		p.Equipement.Torse = nomObjet
		bonusPV = 25
	case "pieds":
		ancien = p.Equipement.Pieds
		p.Equipement.Pieds = nomObjet
		bonusPV = 15
	case "arme":
		ancien = p.Equipement.Arme
		p.Equipement.Arme = nomObjet
		bonusPV = 0
	default:
		fmt.Println("Section d'équipement inconnue.")
		return
	}

	p.PointsDeVieMaximum += bonusPV
	fmt.Printf("%s équipé ! (+%d PV max)\n", nomObjet, bonusPV)

	// si un objet était déjà équipé à cette section, il revient dans l'inventaire
	if ancien != "" {
		p.AddInventory(ancien)
		fmt.Println(ancien, "retourne dans ton inventaire.")
	}
}

// Bonus d'attaque apporté par l'arme équipée.
func (p Personne) BonusAttaqueEquipement() int {
	switch p.Equipement.Arme {
	case "Dague rouillée":
		return 3
	case "Épée d'acier":
		return 7
	case "Arc du chasseur":
		return 5
	default:
		return 0
	}
}

// TACHE 18 : augmenter la capacité de l'inventaire (max 3 fois)
func (p *Personne) UpgradeInventorySlot() {
	if p.AmeliorationsInventaire >= 3 {
		fmt.Println("Tu as déjà utilisé les 3 augmentations d'inventaire possibles.")
		return
	}
	p.AmeliorationsInventaire++
	fmt.Println("Ton inventaire peut maintenant contenir", p.TailleMaxInventaireActuelle(), "objets.")
}

// TACHE 11 : reformate un nom saisi -> première lettre majuscule, le reste en minuscule
func FormaterNom(nom string) string {
	nom = strings.ToLower(nom)
	if len(nom) == 0 {
		return nom
	}
	return strings.ToUpper(nom[:1]) + nom[1:]
}

// TACHE 11 : vérifie que le nom ne contient que des lettres
func NomValide(nom string) bool {
	if len(nom) == 0 {
		return false
	}
	for _, lettre := range nom {
		if !unicode.IsLetter(lettre) {
			return false
		}
	}
	return true
}

// TACHE 11 : statistiques de départ selon la classe choisie
func StatsDeBase(classe string) (pvMax int, ok bool) {
	switch classe {
	case "Humain":
		return 100, true
	case "Elfe":
		return 80, true
	case "Nain":
		return 120, true
	default:
		return 0, false
	}
}
