package structures

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lecteurDonjon = bufio.NewReader(os.Stdin)

func lireDonjon() string {
	texte, _ := lecteurDonjon.ReadString('\n')
	return strings.TrimSpace(texte)
}

// BONUS : nombre de niveaux disponibles dans le donjon
const NombreNiveauxDonjon = 5

// Renvoie un monstre de base tiré au hasard, dont les stats sont mises à
// l'échelle selon le niveau de donjon choisi (1 = normal, 5 = très difficile)
func genererMonstreNiveau(niveau int) Monster {
	m := ChoisirMonsterAleatoire()

	// chaque niveau au-dessus de 1 ajoute 50% de PV/attaque/XP en plus
	multiplicateur := 1.0 + float64(niveau-1)*0.5

	m.PointsDeVieMaximum = int(float64(m.PointsDeVieMaximum) * multiplicateur)
	m.PointsDeVieActuels = m.PointsDeVieMaximum
	m.PointsAttaque = int(float64(m.PointsAttaque) * multiplicateur)
	m.ExperienceOffert = int(float64(m.ExperienceOffert) * multiplicateur)

	// le dernier niveau est un boss : nom marqué, initiative élevée
	if niveau == NombreNiveauxDonjon {
		m.Nom = m.Nom + " (BOSS)"
		m.Initiative += 5
	}

	return m
}

// BONUS : menu de sélection du niveau de donjon
func LancerDonjon(p *Personne) {
	for {
		fmt.Println("\n===== DONJON =====")
		for niveau := 1; niveau <= NombreNiveauxDonjon; niveau++ {
			label := fmt.Sprintf("Niveau %d", niveau)
			if niveau == NombreNiveauxDonjon {
				label += " (BOSS)"
			}
			fmt.Printf("%d. %s\n", niveau, label)
		}
		fmt.Println("0. Retour")
		fmt.Print("Ton choix : ")

		choix := lireDonjon()
		if choix == "0" {
			return
		}

		niveau, err := strconv.Atoi(choix)
		if err != nil || niveau < 1 || niveau > NombreNiveauxDonjon {
			fmt.Println("Choix invalide.")
			continue
		}

		monstre := genererMonstreNiveau(niveau)
		gagne := DeroulerCombat(p, monstre)

		if gagne && niveau == NombreNiveauxDonjon {
			fmt.Println("\n*** Tu as vaincu le boss du donjon ! Bravo ! ***")
		}
	}
}
