package structures

import "fmt"

type Personne struct {
	Nom                string
	Classe             string
	Niveau             int
	PointsDeVieMaximum int
	PointsDeVieActuels int
	Inventaire         []string
}

func InitCharacter(nom string, classe string, niveau int, pvMax int, pvActuels int, inventaire []string) Personne {
	return Personne{
		Nom:                nom,
		Classe:             classe,
		Niveau:             niveau,
		PointsDeVieMaximum: pvMax,
		PointsDeVieActuels: pvActuels,
		Inventaire:         inventaire,
	}
}

func (p Personne) DisplayInfo() {
	fmt.Println("=== Informations du personnage ===")
	fmt.Println("Nom     :", p.Nom)
	fmt.Println("Classe  :", p.Classe)
	fmt.Println("Niveau  :", p.Niveau)
	fmt.Printf("PV      : %d / %d\n", p.PointsDeVieActuels, p.PointsDeVieMaximum)
	fmt.Println("Inventaire :", p.Inventaire)
}