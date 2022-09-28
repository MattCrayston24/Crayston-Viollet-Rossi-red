package Dossier_Package

type Monstre struct {
	nom                 string
	point_de_vie_max    int
	point_de_vie_actuel int
	points_d_attaque    int
	initiative          int
	Experience          int
	element             string
}

func (m *Monstre) InitMonstre() {
	m.nom = "Gobelin d'entrainement"
	m.point_de_vie_max = 40
	m.point_de_vie_actuel = m.point_de_vie_max
	m.points_d_attaque = 5
	m.initiative = 10
	m.element = "Eau"
}

func (m *Monstre) gobelin_parterne(nb_tour int) {
	if nb_tour%3 == 1 {
		m.points_d_attaque *= 2
	} else {
		m.points_d_attaque = 5
	}

}
