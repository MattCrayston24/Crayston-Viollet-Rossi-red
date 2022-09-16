package Dossier_Package

type Monstre struct {
	nom                 string
	point_de_vie_max    int
	point_de_vie_actuel int
	points_d_attaque    int
}

func (m *Monstre) InitMonstre() {
	m.nom = "Gobelin d'entrainement"
	m.point_de_vie_max = 40
	m.point_de_vie_actuel = m.point_de_vie_max
	m.points_d_attaque = 5
}
