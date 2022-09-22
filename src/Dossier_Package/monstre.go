package Dossier_Package

import "fmt"

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

func (m *Monstre) gobelin_parterne(nb_tour int) {
	if nb_tour%3 == 1 {
		m.points_d_attaque *= 2
	} else {
		m.points_d_attaque = 5
	}

}

func (p *Personnage) ChatTurn(m *Monstre) {
	var choix_inventaire int
	var choix_menu int
	for p.point_de_vie_actuel != 0 || m.point_de_vie_actuel != 0 {
		fmt.Print("Si vous voulez attaquer taper 1 \n Si vous voulez accéder à l'inventaire taper 2")
		fmt.Scan(&choix_menu)
		switch choix_menu {
		case 1:
			p.Menu_attaque()
			m.point_de_vie_actuel -= p.points_attaque
			fmt.Print("Gobelin d'entrainement PV :", m.point_de_vie_actuel, "/", m.point_de_vie_max)
		case 2:
			fmt.Print(p.inventaire, "Si vous taper un chiffre superieur a 2 vous quitterez l'inventaire \n Si vous taper 1 vous pourrez utiliser une potion de vie si vous en avez une \n si vous taper 2 vous pourrez alors utiliser une potion de poison si vous en avez une")
			fmt.Scan(&choix_inventaire)
			if choix_inventaire == 1 {
				p.TakePot()
				p.removeInventory("Potion de Vie")
			} else if choix_inventaire == 2 {
				p.PoisonPot(m)
				p.removeInventory("Potion de poison")
			} else {
				fmt.Print("Vous quittez l'inventaire !")
			}
		default:
			fmt.Print("Vous avez entré une valeur éroné")
		}
	}
}

func (p *Personnage) trainingFight() {
	var m1 *Monstre
	m1.InitMonstre()
	var nbtours int
	for i := 0; p.point_de_vie_actuel <= 0 || m1.point_de_vie_actuel <= 0; i++ {
		if i%2 == 0 {
			p.ChatTurn(m1)
			nbtours++
		} else {
			m1.gobelin_parterne(nbtours)
			p.point_de_vie_actuel -= m1.points_d_attaque
			nbtours++
		}
		fmt.Println(nbtours)
	}
}

func (p *Personnage) Menu_attaque() {
	var choix int

	fmt.Println("choisisser le numéro de l'attque que vous voulez effectuer \n1: Basique \n2:Skill")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		p.points_attaque = 5
		fmt.Print("Vous utilisez attaque basique et infligé 5 points de dégats")
	case 2:

	default:
		fmt.Println("Vous avez entré une valeur éroné")
	}

}
