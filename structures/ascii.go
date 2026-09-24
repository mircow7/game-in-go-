package structures
 
import (
	"fmt"
	"strings"
)
 
// ============================================================
// BONUS : couleurs terminal (codes ANSI)
// ============================================================
 
const (
	couleurReset  = "\033[0m"
	couleurVert   = "\033[32m"
	couleurJaune  = "\033[33m"
	couleurRouge  = "\033[31m"
	couleurCyan   = "\033[36m"
	couleurViolet = "\033[35m"
)
 
// ============================================================
// BONUS : barre de vie visuelle
// ============================================================
 
// BarrePVCouleur affiche une barre du type [##########----------] 40/100
// colorée en vert, jaune ou rouge selon le pourcentage de vie restant
func BarrePVCouleur(actuel int, max int) string {
	if max <= 0 {
		max = 1
	}
	if actuel < 0 {
		actuel = 0
	}
 
	largeur := 20
	rempli := int(float64(actuel) / float64(max) * float64(largeur))
	if rempli > largeur {
		rempli = largeur
	}
	vide := largeur - rempli
 
	barre := strings.Repeat("#", rempli) + strings.Repeat("-", vide)
 
	pourcentage := float64(actuel) / float64(max)
	couleur := couleurVert
	if pourcentage < 0.3 {
		couleur = couleurRouge
	} else if pourcentage < 0.6 {
		couleur = couleurJaune
	}
 
	return fmt.Sprintf("%s[%s]%s %d/%d", couleur, barre, couleurReset, actuel, max)
}
 
// BarreSimple : même principe mais sans couleur, pour le mana par exemple
func BarreSimple(actuel int, max int) string {
	if max <= 0 {
		max = 1
	}
	if actuel < 0 {
		actuel = 0
	}
 
	largeur := 20
	rempli := int(float64(actuel) / float64(max) * float64(largeur))
	if rempli > largeur {
		rempli = largeur
	}
	vide := largeur - rempli
 
	barre := strings.Repeat("#", rempli) + strings.Repeat("-", vide)
	return fmt.Sprintf("%s[%s]%s %d/%d", couleurCyan, barre, couleurReset, actuel, max)
}
 
// ============================================================
// BONUS : écran titre
// ============================================================
 
const titreASCII = `
  _____  _____   ____      _ ______ _____   _____  ______ _____
 |  __ \|  __ \ / __ \    | |  ____|_   _| |  __ \|  ____|  __ \
 | |__) | |__) | |  | |   | | |__    | |   | |__) | |__  | |  | |
 |  ___/|  _  /| |  | |   | |  __|   | |   |  _  /|  __| | |  | |
 | |    | | \ \| |__| |   | | |____ _| |_  | | \ \| |____| |__| |
 |_|    |_|  \_\\____/    |_|______|_____| |_|  \_\______|_____/
`
 
func AfficherTitre() {
	fmt.Println(couleurViolet + titreASCII + couleurReset)
}
 
// ============================================================
// BONUS : victoire / défaite
// ============================================================
 
const victoireASCII = `
   \\o/    VICTOIRE !
    |
   / \
`
 
const defaiteASCII = `
    x_x    K.O. !
   /|||\
`
 
func AfficherVictoire() {
	fmt.Println(couleurVert + victoireASCII + couleurReset)
}
 
func AfficherDefaite() {
	fmt.Println(couleurRouge + defaiteASCII + couleurReset)
}
 
// ============================================================
// BONUS : visuels des monstres
// ============================================================
 
const artGobelin = `
    ,     ,
   (\____/)
    (_oo_)
     (O)
   __||||__
`
 
const artLoup = `
   /\   /\
  {  ' - '  }
   \  ~  /
   / '---' \
`
 
const artTroll = `
     .-""""-.
    /  o  o  \
   |    ~     |
    \  ===  /
   /|       |\
  ' |       | '
`
 
const artGenerique = `
   [ ??? ]
   un monstre
   inconnu
`
 
// ObtenirArtMonstre renvoie l'ASCII art correspondant au nom du monstre.
// Fonctionne aussi si le nom porte le suffixe " (BOSS)" ajouté par le donjon.
func ObtenirArtMonstre(nom string) string {
	nomBase := strings.TrimSuffix(nom, " (BOSS)")
 
	switch nomBase {
	case "Gobelin d'entrainement":
		return artGobelin
	case "Loup sauvage":
		return artLoup
	case "Troll des marais":
		return artTroll
	default:
		return artGenerique
	}
}
 