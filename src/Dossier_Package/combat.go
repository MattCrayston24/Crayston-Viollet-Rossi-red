package Dossier_Package

import (
	"fmt"
)	

func (p *Personnage) ChatTurn(m *Monstre) {
	var choix_inventaire string
	var choix_menu string
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                Combat              |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	fmt.Println("|             1.Attaquer             |")
	fmt.Println("|                                    |")
	fmt.Println("|             2.Inventaire           |")
	fmt.Println("|____________________________________|")
	fmt.Scanln(&choix_menu)
	p.verif_espace(choix_menu)
	switch choix_menu {
	case "1":
		p.Menu_attaque(m)
		m.point_de_vie_actuel -= p.points_attaque
		fmt.Println(m.nom, "PV :", m.point_de_vie_actuel, "/", m.point_de_vie_max)
	case "2":
		fmt.Print(p.inventaire, "Si vous taper un chiffre superieur a 2 vous quitterez l'inventaire \n Si vous taper 1 vous pourrez utiliser une potion de vie si vous en avez une \n si vous taper 2 vous pourrez alors utiliser une potion de poison si vous en avez une")
		fmt.Println("______________________________________")
		fmt.Println("|                                    |")
		fmt.Println("|   1.Potion de vie                  |")
		fmt.Println("|                                    |")
		fmt.Println("|   2.Potion de poison               |")
		fmt.Println("|                                    |")
		fmt.Println("|   Autre. Quitter                   |")
		fmt.Println("|____________________________________|")
		fmt.Scan(&choix_inventaire)
		p.verif_espace(choix_inventaire)
		if choix_inventaire == "1" {
			p.TakePot()
			p.removeInventory("Potion de Vie")
		} else if choix_inventaire == "2" {
			p.PoisonPot(m)
			p.removeInventory("Potion de poison")
		} else {
			fmt.Print("Vous quittez l'inventaire !")
		}
	default:
		fmt.Print("Vous avez entré une valeur éroné")
		p.ChatTurn(m)
	}
}

func (p *Personnage) trainingFight() {
	var m1 *Monstre = new(Monstre)
	m1.InitMonstre(Menu_Choix_Monstre())
	fmt.Println("le monstre a été initialiser en tant que :", m1.nom)
	var nbtours int
	var compt_mort int 
	if p.initiative < m1.initiative {
		nbtours = 1
	}
	for ;p.point_de_vie_actuel >= 0 || m1.point_de_vie_actuel > 0;{
		if nbtours%2 == 0 {
			p.ChatTurn(m1)
			nbtours++
		} else {
			m1.gobelin_parterne(nbtours)
			p.point_de_vie_actuel -= m1.points_d_attaque
			fmt.Println("le monstre vous attauque vous perdez", m1.points_d_attaque, "il vous reste", p.point_de_vie_actuel, "/", p.point_de_vie_maximum)
			nbtours++
		}
		fmt.Println("nombre de tour :", nbtours)
		if p.point_de_vie_actuel <= 0 && compt_mort >= 1 {
			fmt.Println("t'es mort cheh")
			break
		}else if m1.point_de_vie_actuel <= 0{
			fmt.Println("t'es mort ")
			break
		}
		if p.wasted() && compt_mort!=1 {
			p.point_de_vie_actuel = p.point_de_vie_maximum /2
			fmt.Println("votre personnage est mort mais vous revenez a la vie avec :",p.point_de_vie_actuel,"de PV ")
			compt_mort = 1 
		}
	
	}
	if p.point_de_vie_actuel > 0 {
		p.addExp(m1.Experience)
		fmt.Println("l'experience a été ajouté")
		for j := 0; j < len(m1.drop); j++ {
			p.addInventory(m1.drop[j])
			fmt.Println(m1.drop[j], " a été ajouter a votre inventaire")
		}
		p.ajout_monnaie(m1.monnaie)
		fmt.Println(m1.monnaie, "a été ajoutée a votre monnaie")
	}
	p.Menu()
}

func (p *Personnage) Menu_attaque(m *Monstre) {
	var choix string

	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|               Menu                 |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	fmt.Println("|             1.Attaque Basique      |")
	fmt.Println("|                                    |")
	fmt.Println("|             2.Skill                |")
	fmt.Println("|____________________________________|")
	fmt.Scan(&choix)
	p.verif_espace(choix)

	switch choix {
	case "1":
		p.points_attaque = p.attaque_base
		fmt.Println("Vous utilisez attaque basique et infligé ", p.points_attaque, "points de dégats")
	case "2":
		p.Menu_skill(m)
	default:
		fmt.Println("Vous avez entré une valeur éroné")
	}

}