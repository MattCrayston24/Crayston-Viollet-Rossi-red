package Dossier_Package

import "fmt"

func (p *Personnage) ChatTurn(m *Monstre){
	var choix_inventaire int
	var choix_menu int
	for p.point_de_vie_actuel != 0 || m.point_de_vie_actuel != 0 {
		fmt.Print("Si vous voulez attaquer taper 1 \n Si vous voulez accéder à l'inventaire taper 2")
		fmt.Scan(&choix_menu)
		switch choix_menu {
		case 1:
			fmt.Print("Vous utilisez attaque basique et infligé 5 points de dégats")	
			m.point_de_vie_actuel -= 5
			fmt.Print("Gobelin d'entrainement PV :",m.point_de_vie_actuel,"/",m.point_de_vie_max)
		case 2:
			fmt.Print(p.inventaire, "Si vous taper un chiffre superieur a 2 vous quitterez l'inventaire \n Si vous taper 1 vous pourrez utiliser une potion de vie si vous en avez une \n si vous taper 2 vous pourrez alors utiliser une potion de poison si vous en avez une")
			fmt.Scan(&choix_inventaire)
			if choix_inventaire == 1 {
				p.TakePot()
				p.removeInventory("Potion de Vie")
			}else if choix_inventaire == 2 {
				p.PoisonPot(m)
				p.removeInventory("Potion de poison")
			}else {
				fmt.Print("Vous quittez l'inventaire !")
			}
		default :
			fmt.Print("Vous avez entré une valeur éroné")	
		}
	}
}

