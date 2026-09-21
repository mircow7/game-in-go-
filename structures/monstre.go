package structures
 
import (
	"fmt"
	"math/rand"
)
 
// TACHE 19 : structure du monstre
// MISSION 1 : + Initiative   MISSION 2 : + ExperienceOffert   BONUS : + Butin
type Monster struct {
	Nom                string
	PointsDeVieMaximum int
	PointsDeVieActuels int
	PointsAttaque      int
	Initiative         int
	ExperienceOffert   int
	Butin              []string
}
 
// TACHE 19 suite : initialisation du Gobelin d'entrainement
func InitGoblin() Monster {
	return Monster{
		Nom:                "Gobelin d'entrainement",
		PointsDeVieMaximum: 40,
		PointsDeVieActuels: 40,
		PointsAttaque:      5,
		Initiative:         RandomInitiative(),
		ExperienceOffert:   30,
		Butin:              []string{"Plume de Corbeau", "Cuir de Sanglier", "Potion de vie"},
	}
}
 
// BONUS : deux monstres supplémentaires pour varier l'entrainement
func InitLoup() Monster {
	return Monster{
		Nom:                "Loup sauvage",
		PointsDeVieMaximum: 30,
		PointsDeVieActuels: 30,
		PointsAttaque:      7,
		Initiative:         RandomInitiative(),
		ExperienceOffert:   25,
		Butin:              []string{"Fourrure de Loup", "Fourrure de Loup", "Potion de vie"},
	}
}
 
func InitTroll() Monster {
	return Monster{
		Nom:                "Troll des marais",
		PointsDeVieMaximum: 60,
		PointsDeVieActuels: 60,
		PointsAttaque:      4,
		Initiative:         RandomInitiative(),
		ExperienceOffert:   45,
		Butin:              []string{"Peau de Troll", "Potion de mana", "AK-47"},
	}
}
 
// BONUS : tire un monstre au hasard parmi ceux disponibles
func ChoisirMonsterAleatoire() Monster {
	generateurs := []func() Monster{InitGoblin, InitLoup, InitTroll}
	index := rand.Intn(len(generateurs))
	return generateurs[index]()
}
 
func (m Monster) AfficherPV() {
	fmt.Printf("%s : PV %d / %d\n", m.Nom, m.PointsDeVieActuels, m.PointsDeVieMaximum)
}
 
func (m Monster) EstMort() bool {
	return m.PointsDeVieActuels <= 0
}
 
// TACHE 20 : pattern d'attaque du gobelin
// tous les 3 tours, il inflige 200% de son attaque au lieu de 100%
func (m *Monster) GoblinPattern(p *Personne, tour int) {
	degats := m.PointsAttaque
	if tour%3 == 0 {
		degats = m.PointsAttaque * 2
	}
 
	p.PointsDeVieActuels -= degats
	if p.PointsDeVieActuels < 0 {
		p.PointsDeVieActuels = 0
	}
 
	fmt.Printf("%s inflige à %s %d de dégâts\n", m.Nom, p.Nom, degats)
	p.AfficherPV()
}
 