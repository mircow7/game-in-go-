package structures

import (
	"fmt"
	"math/rand"
)

// Monster représente un ennemi. SpriteID est utilisé par la version graphique.
type Monster struct {
	Nom                string
	PointsDeVieMaximum int
	PointsDeVieActuels int
	PointsAttaque      int
	Initiative         int
	ExperienceOffert   int
	Butin              []string
	SpriteID           string
	OrRecompense       int
}

func InitGoblin() Monster {
	return Monster{
		Nom: "Gobelin d'entrainement", PointsDeVieMaximum: 40, PointsDeVieActuels: 40,
		PointsAttaque: 5, Initiative: RandomInitiative(), ExperienceOffert: 30,
		Butin:    []string{"Plume de Corbeau", "Cuir de Sanglier", "Potion de vie"},
		SpriteID: "pumpkin", OrRecompense: 4,
	}
}

func InitLoup() Monster {
	return Monster{
		Nom: "Loup sauvage", PointsDeVieMaximum: 30, PointsDeVieActuels: 30,
		PointsAttaque: 7, Initiative: RandomInitiative(), ExperienceOffert: 25,
		Butin:    []string{"Fourrure de Loup", "Fourrure de Loup", "Potion de vie"},
		SpriteID: "zombie", OrRecompense: 5,
	}
}

func InitTroll() Monster {
	return Monster{
		Nom: "Troll des marais", PointsDeVieMaximum: 60, PointsDeVieActuels: 60,
		PointsAttaque: 4, Initiative: RandomInitiative(), ExperienceOffert: 45,
		Butin:    []string{"Peau de Troll", "Potion de mana"},
		SpriteID: "knight", OrRecompense: 8,
	}
}

// Nouvelles créatures graphiques du donjon.
func InitPumpkin() Monster {
	return Monster{
		Nom: "Citrouille maudite", PointsDeVieMaximum: 50, PointsDeVieActuels: 50,
		PointsAttaque: 8, Initiative: RandomInitiative(), ExperienceOffert: 40,
		Butin:    []string{"Potion de vie", "Potion de mana", "Cuir de Sanglier"},
		SpriteID: "pumpkin", OrRecompense: 10,
	}
}

func InitZombie() Monster {
	return Monster{
		Nom: "Zombie des cryptes", PointsDeVieMaximum: 65, PointsDeVieActuels: 65,
		PointsAttaque: 9, Initiative: RandomInitiative(), ExperienceOffert: 55,
		Butin:    []string{"Potion de vie", "Peau de Troll", "Potion de poison"},
		SpriteID: "zombie", OrRecompense: 12,
	}
}

func InitChevalier() Monster {
	return Monster{
		Nom: "Chevalier maudit", PointsDeVieMaximum: 90, PointsDeVieActuels: 90,
		PointsAttaque: 11, Initiative: RandomInitiative(), ExperienceOffert: 80,
		Butin:    []string{"Potion de mana", "Livre de Sort : Boule de Feu"},
		SpriteID: "knight", OrRecompense: 20,
	}
}

func ChoisirMonsterAleatoire() Monster {
	generateurs := []func() Monster{
		InitGoblin, InitLoup, InitTroll,
		InitPumpkin, InitZombie, InitChevalier,
	}
	return generateurs[rand.Intn(len(generateurs))]()
}

func (m Monster) AfficherPV() {
	fmt.Printf("%s : PV %d / %d\n", m.Nom, m.PointsDeVieActuels, m.PointsDeVieMaximum)
}

func (m Monster) EstMort() bool {
	return m.PointsDeVieActuels <= 0
}

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
