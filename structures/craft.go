package structures

import "fmt"

// TACHE 15 : recette = liste d'ingrédients nécessaires
type Ingredient struct {
	Nom      string
	Quantite int
}

type Recette struct {
	NomObjet    string
	Prix        int
	Ingredients []Ingredient
	Section     string // "tete", "torse" ou "pieds" -> utilisé à l'équipement (TACHE 17)
}

var Recettes = map[string]Recette{
	"Chapeau de l'aventurier": {
		NomObjet: "Chapeau de l'aventurier", Prix: 5, Section: "tete",
		Ingredients: []Ingredient{{"Plume de Corbeau", 1}, {"Cuir de Sanglier", 1}},
	},
	"Tunique de l'aventurier": {
		NomObjet: "Tunique de l'aventurier", Prix: 5, Section: "torse",
		Ingredients: []Ingredient{{"Fourrure de Loup", 2}, {"Peau de Troll", 1}},
	},
	"Bottes de l'aventurier": {
		NomObjet: "Bottes de l'aventurier", Prix: 5, Section: "pieds",
		Ingredients: []Ingredient{{"Fourrure de Loup", 1}, {"Cuir de Sanglier", 1}},
	},
	"Dague rouillée": {
		NomObjet: "Dague rouillée", Prix: 8, Section: "arme",
		Ingredients: []Ingredient{{"Cuir de Sanglier", 1}, {"Fer", 2}},
	},
	"Épée d'acier": {
		NomObjet: "Épée d'acier", Prix: 20, Section: "arme",
		Ingredients: []Ingredient{{"Fer", 3}, {"Peau de Troll", 1}},
	},
	"Arc du chasseur": {
		NomObjet: "Arc du chasseur", Prix: 14, Section: "arme",
		Ingredients: []Ingredient{{"Bois", 2}, {"Cuir de Sanglier", 2}},
	},
	"Potion supérieure": {
		NomObjet: "Potion supérieure", Prix: 12, Section: "objet",
		Ingredients: []Ingredient{{"Herbe médicinale", 2}, {"Fourrure de Loup", 1}},
	},
	"Elixir de mana": {
		NomObjet: "Elixir de mana", Prix: 12, Section: "objet",
		Ingredients: []Ingredient{{"Herbe médicinale", 2}, {"Plume de Corbeau", 1}},
	},
}

// Compte combien d'exemplaires d'un objet le personnage possède
func (p Personne) CompteItem(nom string) int {
	compte := 0
	for _, item := range p.Inventaire {
		if item == nom {
			compte++
		}
	}
	return compte
}

// Vérifie que le personnage a bien tous les ingrédients de la recette
func (p Personne) PeutFabriquer(r Recette) bool {
	for _, ing := range r.Ingredients {
		if p.CompteItem(ing.Nom) < ing.Quantite {
			return false
		}
	}
	return true
}

// TACHE 15 : fabrique un objet si l'argent et les ressources sont suffisants
func (p *Personne) Fabriquer(nomObjet string) {
	recette, existe := Recettes[nomObjet]
	if !existe {
		fmt.Println("Cet objet ne peut pas être fabriqué.")
		return
	}

	if p.InventairePlein() {
		fmt.Println("Ton inventaire est plein, impossible de fabriquer.")
		return
	}
	if p.Argent < recette.Prix {
		fmt.Println("Tu n'as pas assez d'or pour fabriquer cet objet.")
		return
	}
	if !p.PeutFabriquer(recette) {
		fmt.Println("Il te manque des ressources pour fabriquer :", nomObjet)
		for _, ing := range recette.Ingredients {
			fmt.Printf("  - %s : %d/%d\n", ing.Nom, p.CompteItem(ing.Nom), ing.Quantite)
		}
		return
	}

	// on retire les ingrédients (autant de fois que nécessaire)
	for _, ing := range recette.Ingredients {
		for i := 0; i < ing.Quantite; i++ {
			p.RemoveInventory(ing.Nom)
		}
	}
	p.Argent -= recette.Prix
	p.AddInventory(nomObjet)
	fmt.Println("Tu as fabriqué :", nomObjet)
}
