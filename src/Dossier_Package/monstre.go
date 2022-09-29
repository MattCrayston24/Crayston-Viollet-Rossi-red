package Dossier_Package

import "fmt"

type Monstre struct {
	nom                 string
	point_de_vie_max    int
	point_de_vie_actuel int
	points_d_attaque    int
	initiative          int
	Experience          int
	element             string
	drop                []string
	monnaie             int
}

func (m *Monstre) InitMonstre(str string) {
	if str == "Gobelin" {
		m.nom = "Gobelin"
		m.point_de_vie_max = 40
		m.point_de_vie_actuel = m.point_de_vie_max
		m.points_d_attaque = 5
		m.initiative = 10
		m.element = "Eau"
		m.drop = []string{"Cuir", "Cuir", "Cuir"}
		m.monnaie = 30
<<<<<<< HEAD
=======
		m.Experience = 80
>>>>>>> main
	} else if str == "Kobolt" {
		m.nom = "Kobolt"
		m.point_de_vie_max = 60
		m.point_de_vie_actuel = m.point_de_vie_max
		m.points_d_attaque = 8
		m.initiative = 15
		m.element = "Terre"
		m.drop = []string{"Fer", "Fer", "Fer"}
		m.monnaie = 50
<<<<<<< HEAD
=======
		m.Experience = 50
>>>>>>> main
	} else if str == "Orc" {
		m.nom = "Orc"
		m.point_de_vie_max = 80
		m.point_de_vie_actuel = m.point_de_vie_max
		m.points_d_attaque = 10
		m.initiative = 20
		m.element = "Feu"
		m.drop = []string{"Mithril", "Mithril", "Mithril"}
		m.monnaie = 80
<<<<<<< HEAD
=======
		m.Experience = 80
>>>>>>> main
	}
}

func (m *Monstre) gobelin_parterne(nb_tour int) {
	if nb_tour%3 == 1 {
		m.points_d_attaque *= 2
	} else {
		m.points_d_attaque = 5
	}

}

func Menu_Choix_Monstre() string {
<<<<<<< HEAD
	var choix string
=======
	var choix int
>>>>>>> main
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|               Menu                 |")
	fmt.Println("|               Monstre              |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	fmt.Println("|             1.Gobelin              |")
	fmt.Println("|                                    |")
	fmt.Println("|             2.Kobolt               |")
	fmt.Println("|                                    |")
	fmt.Println("|             3. Orc                 |")
	fmt.Println("|                                    |")
	fmt.Println("|____________________________________|")
	fmt.Scan(&choix)
<<<<<<< HEAD
	switch choix {
	case "1":
		return "Gobelin"
	case "2":
		return "Kobolt"
	case "3":
		return "Orc"	
=======

	switch choix {
	case 1:
		return "Gobelin"
	case 2:
		return "Kobolt"
	case 3:
		return "Orc"
>>>>>>> main
	}
	return "aucun monstre choisi"
}
