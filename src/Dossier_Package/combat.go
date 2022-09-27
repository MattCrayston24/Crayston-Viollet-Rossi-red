package Dossier_Package

import "fmt"

func (p *Personnage) ChatTurn(m *Monstre) {
	var choix_inventaire int
	var choix_menu int
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
	p.points_attaque = 3 * p.niveau
}

func (p *Personnage) trainingFight() {
	var m1 *Monstre
	m1.InitMonstre()
	var nbtours int
	i := 0
	if p.initiative < m1.initiative {
		i = 1
	}
	for ; p.point_de_vie_actuel <= 0 || m1.point_de_vie_actuel <= 0; i++ {
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
	if p.point_de_vie_actuel < 0 {
		p.addExp(m1.Experience)
	}
}

func (p *Personnage) Menu_attaque() {
	var choix int

	fmt.Println("choisisser le numéro de l'attque que vous voulez effectuer \n1: Basique \n2:Skill")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		fmt.Print("Vous utilisez attaque basique et infligé ", p.points_attaque, "points de dégats")
	case 2:
		p.Menu_skill()
		fmt.Print("Vous utilisez un skill et infligé ", p.points_attaque, "points de dégats")
	default:
		fmt.Println("Vous avez entré une valeur éroné")
	}

}