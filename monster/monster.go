package structures

import "fmt"

type Monster struct {
	Nom                string
	PointsDeVieMaximum int
	PointsDeVieActuels int
	PointsAttaque      int
}


func InitGoblin() Monster {
	return Monster{
		Nom:                "Gobelin d'entrainement",
		PointsDeVieMaximum: 40,
		PointsDeVieActuels: 40,
		PointsAttaque:      5,
	}
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
