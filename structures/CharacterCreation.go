package structures

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var lecteurCreation = bufio.NewReader(os.Stdin)

func lireCreation() string {
	texte, _ := lecteurCreation.ReadString('\n')
	return strings.TrimSpace(texte)
}

// TACHE 11 : laisse le joueur créer son personnage lui-même
// Remplace l'appel direct à InitCharacter fait dans la TACHE 2
func CharacterCreation() Personne {
	// --- Choix du nom ---
	var nom string
	for {
		fmt.Print("Choisis le nom de ton personnage (lettres uniquement) : ")
		saisie := lireCreation()

		if !NomValide(saisie) {
			fmt.Println("Nom invalide, utilise uniquement des lettres.")
			continue
		}

		nom = FormaterNom(saisie)
		break
	}

	// --- Choix de la classe ---
	var classe string
	var pvMax int
	for {
		fmt.Println("Choisis ta classe : Humain, Elfe ou Nain")
		fmt.Print("Ton choix : ")
		saisie := lireCreation()

		pv, ok := StatsDeBase(saisie)
		if !ok {
			fmt.Println("Classe invalide, choisis Humain, Elfe ou Nain.")
			continue
		}

		classe = saisie
		pvMax = pv
		break
	}

	// --- TACHE 11 suite : points de vie actuels = 50% du max, niveau 1 ---
	pvActuels := pvMax / 2

	perso := InitCharacter(nom, classe, 1, pvMax, pvActuels, []string{})

	// MISSION 1, 2, 4 : valeurs de départ
	perso.Initiative = RandomInitiative()
	perso.ManaMax = 30
	perso.Mana = perso.ManaMax
	perso.ExperienceActuelle = 0
	perso.ExperienceMax = ExperienceRequiseBase

	fmt.Printf("\nBienvenue %s le %s !\n", perso.Nom, perso.Classe)
	perso.AfficherPV()

	return perso
}
