package Dossier_Package

import (
	"fmt"
	"os"
)

// Fonction pour définir le menu :
func (p Personnage) Menu() {
	var choix_menu int
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|               Menu                 |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	fmt.Println("|             1.Inventaire           |")
	fmt.Println("|                                    |")
	fmt.Println("|             2.Stat Perso           |")
	fmt.Println("|                                    |")
	fmt.Println("|             3.Forgeron             |")
	fmt.Println("|                                    |")
	fmt.Println("|             4.Marchand             |")
	fmt.Println("|                                    |")
	fmt.Println("|             5.Combat Entrainement  |")
	fmt.Println("|                                    |")
	fmt.Println("|             6.Menu Skill           |")
	fmt.Println("|                                    |")
	fmt.Println("|             7.Quitter              |")
	fmt.Println("|____________________________________|")
	fmt.Scan(&choix_menu)

	switch choix_menu {
	case 1:
		p.AccèsInventaire()
	case 2:
		p.Afficher_info()
	case 3:
		p.Menu_Fogeron()
	case 4:
		p.Marchand()
	case 5:
		p.trainingFight()
	case 6:
		p.Menu_skill_choix()
	case 7:
		os.Exit(3)
	default:
		fmt.Println("Vous n'avez pas selectionner un reponse valide")
		p.Menu()
	}
}

func (p Personnage) Menu_Inventaire() {
	var choix int
	fmt.Println(p.inventaire)
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                Inventaire          |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	fmt.Println("|             1.Prendre Potion       |")
	fmt.Println("|                de vie              |")
	fmt.Println("|                                    |")
	fmt.Println("|             2.Mettre un            |")
	fmt.Println("|               équipement           |")
	fmt.Println("|                                    |")
	fmt.Println("|             3.Retour au            |")
	fmt.Println("|               Précédent            |")
	fmt.Println("|____________________________________|")
	fmt.Scan(&choix)
	for {
		switch choix {
		case 1:
			p.TakePot()
			p.Menu()
		case 2:
			p.Menu_Equipement()
		case 3:
			p.Menu()
		default:
			fmt.Println("Vous n'avez pas ecrit un reponse valide ")
			p.Menu_Inventaire()
		}
	}
}

func (p Personnage) Menu_Equipement() {
	for {
		var choix int
		fmt.Println("______________________________________")
		fmt.Println("|                                    |")
		fmt.Println("|                Menu                |")
		fmt.Println("|                Equipement          |")
		fmt.Println("|____________________________________|")
		fmt.Println("|                                    |")
		fmt.Println("|             1.Mettre un            |")
		fmt.Println("|               casque               |")
		fmt.Println("|                                    |")
		fmt.Println("|             2.Mettre un            |")
		fmt.Println("|               plastron             |")
		fmt.Println("|                                    |")
		fmt.Println("|             3.Mettre une           |")
		fmt.Println("|               paire de             |")
		fmt.Println("|               bottes               |")
		fmt.Println("|                                    |")
		fmt.Println("|             4.équiper une          |")
		fmt.Println("|               arme                 |")
		fmt.Println("|                                    |")
		fmt.Println("|            5.Retour au menu        |")
		fmt.Println("|              précédent             |")
		fmt.Println("|____________________________________|")
		fmt.Scan(&choix)
		switch choix {
		case 1:
			p.Menu_casque()
		case 2:
			p.Menu_plastron()
		case 3:
			p.Menu_bottes()
		case 4:
			if p.verif_classe() == "archer" {
				p.Menu_armes_archer()
			} else if p.verif_classe() == "mage" {
				p.Menu_armes_mage()
			} else if p.verif_classe() == "épéiste" {
				p.Menu_armes_épéiste()
			} else {
				fmt.Println("error")
			}
		case 5:
			p.AccèsInventaire()
		default:
			fmt.Println("Vous n'avez pas ecrit un reponse valide ")
			p.Menu_Equipement()
		}
	}
}

func (p Personnage) Menu_casque() {
	var e1 Equipement
	e1.Init_List()
	var choix int
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                Casque              |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	for i := 0; i < len(e1.liste_casque); i++ {
		fmt.Println("|             ", i+1, ".", e1.liste_casque[i], "   |")
		fmt.Println("|                                    |")
	}
	fmt.Println("|____________________________________|")
	fmt.Scan(&choix)
	switch choix {
	case 1:
		p.Mettre_casque(1, e1)
		p.Menu_Equipement()
	case 2:
		p.Mettre_casque(2, e1)
		p.Menu_Equipement()
	case 3:
		p.Mettre_casque(3, e1)
		p.Menu_Equipement()
	case 4:
		p.Menu_Equipement()
	default:
		fmt.Println("Vous n'avez pas ecrit un reponse valide ")
		p.Menu_casque()

	}

}

func (p Personnage) Menu_plastron() {
	var e1 Equipement
	e1.Init_List()
	var choix int
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                Plastron            |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	for i := 0; i < len(e1.liste_plastron); i++ {
		fmt.Println("|             ", i+1, ".", e1.liste_plastron[i], "   |")
		fmt.Println("|                                    |")
	}
	fmt.Println("|____________________________________|")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		p.Mettre_Plastron(1, e1)
		p.Menu_Equipement()
	case 2:
		p.Mettre_Plastron(2, e1)
		p.Menu_Equipement()
	case 3:
		p.Mettre_Plastron(3, e1)
		p.Menu_Equipement()
	case 4:
		p.Menu_Equipement()
	default:
		fmt.Println("Vous n'avez pas ecrit un reponse valide ")
		p.Menu_plastron()

	}
}

func (p Personnage) Menu_bottes() {
	var e1 Equipement
	e1.Init_List()
	var choix int
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                Jambière            |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	for i := 0; i < len(e1.liste_bottes); i++ {
		fmt.Println("|             ", i+1, ".", e1.liste_bottes[i], "   |")
		fmt.Println("|                                    |")
	}
	fmt.Println("|____________________________________|")
	fmt.Scan(&choix)
	switch choix {
	case 1:
		p.Mettre_bottes(1, e1)
		p.Menu_Equipement()
	case 2:
		p.Mettre_bottes(2, e1)
		p.Menu_Equipement()
	case 3:
		p.Mettre_bottes(3, e1)
		p.Menu_Equipement()
	case 4:
		p.Menu_Equipement()
	default:
		p.Menu_bottes()

	}
}

func (p *Personnage) Menu_skill(m *Monstre) {
	var choix int
	fmt.Println("Votre réserve de mana est de ", p.mana_actuel, "/", p.mana_maximum)
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                skill               |")
	fmt.Println("|____________________________________|")
	for i := 0; i < len(p.skill.nom); i++ {
		fmt.Println("|             ", i+1, ".", p.skill.nom[i], "   |")
	}
	fmt.Println("|____________________________________|")
	fmt.Scan(&choix)
	if p.mana_actuel >= p.skill.cout_mana[choix-1] {
		p.mana_actuel -= p.skill.cout_mana[choix-1]
		p.points_attaque = p.skill.degat[choix-1]
		if verif_element_faible(p.skill.element[choix-1], m.element) {
			p.points_attaque /= 2
		} else if verif_element_fort(p.skill.element[choix-1], m.element) {
			p.points_attaque *= 2
		}
		m.point_de_vie_actuel -= p.points_attaque
		fmt.Println("Vous utilisez un skill et infligé ", p.points_attaque, "points de dégats")
	} else {
		fmt.Println("Vous n'avez pas assez de mana pour utiliser ce sort")
		p.Menu_attaque(m)
	}
}

func (p *Personnage) Menu_armes_archer() {
	var e1 Equipement
	e1.Init_List()
	var choix int
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                Armes               |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	for i := 0; i < len(e1.liste_armes_archer); i++ {
		fmt.Println("|             ", i+1, ".", e1.liste_armes_archer[i], "   |")
		fmt.Println("|                                    |")
	}
	fmt.Println("|____________________________________|")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		p.Mettre_armes_archer(1, e1)
		p.Menu_Equipement()
	case 2:
		p.Mettre_armes_archer(2, e1)
		p.Menu_Equipement()
	case 3:
		p.Mettre_armes_archer(3, e1)
		p.Menu_Equipement()
	case 4:
		p.Mettre_armes_archer(4, e1)
		p.Menu_Equipement()
	default:
		fmt.Println("Vous n'avez pas ecrit un reponse valide ")
		p.Menu_armes_archer()
	}
}

func (p *Personnage) Menu_armes_mage() {
	var e1 Equipement
	e1.Init_List()
	var choix int
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                Armes               |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	for i := 0; i < len(e1.liste_armes_mage); i++ {
		fmt.Println("|             ", i+1, ".", e1.liste_armes_mage[i], "   |")
		fmt.Println("|                                    |")
	}
	fmt.Println("|____________________________________|")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		p.Mettre_armes_mage(1, e1)
		p.Menu_Equipement()
	case 2:
		p.Mettre_armes_mage(2, e1)
		p.Menu_Equipement()
	case 3:
		p.Mettre_armes_mage(3, e1)
		p.Menu_Equipement()
	case 4:
		p.Mettre_armes_mage(4, e1)
		p.Menu_Equipement()
	default:
		fmt.Println("Vous n'avez pas ecrit un reponse valide ")
		p.Menu_armes_mage()

	}
}

func (p *Personnage) Menu_armes_épéiste() {
	var e1 Equipement
	e1.Init_List()
	var choix int
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                Armes               |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	for i := 0; i < len(e1.liste_armes_épéiste); i++ {
		fmt.Println("|             ", i+1, ".", e1.liste_armes_épéiste[i], "   |")
		fmt.Println("|                                    |")
	}
	fmt.Println("|____________________________________|")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		p.Mettre_armes_épéiste(1, e1)
		p.Menu_Equipement()
	case 2:
		p.Mettre_armes_épéiste(2, e1)
		p.Menu_Equipement()
	case 3:
		p.Mettre_armes_épéiste(3, e1)
		p.Menu_Equipement()
	case 4:
		p.Mettre_armes_épéiste(4, e1)
		p.Menu_Equipement()
	default:
		fmt.Println("Vous n'avez pas ecrit un reponse valide ")
		p.Menu_armes_épéiste()
	}
}
