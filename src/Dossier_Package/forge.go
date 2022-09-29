package Dossier_Package

import "fmt"

//Fonction qui sert de menu pour le forgeron
func (p Personnage) Menu_Fogeron() {
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                Forgeron            |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	fmt.Println("|             1.Equipement           |")
	fmt.Println("|                                    |")
	fmt.Println("|             2.Armes                |")
	fmt.Println("|                                    |")
	fmt.Println("|             3.Quitter              |")
	fmt.Println("|____________________________________|")
	switch p.Scan() {
	case "1":
		p.Menu_Equipement_Creation()
	case "2":
		if p.verif_classe() == "mage" {
			p.Menu_arme_mage_Creation()
		} else if p.verif_classe() == "archer" {
			p.Menu_arme_archer_Creation()
		} else if p.verif_classe() == "épéiste" {
			p.Menu_arme_épéiste_Creation()
		}
	case "3":
		p.Menu()
	default:
		break
	}
}

//Fonction qui permet de fabriquer un objet grace au matériaux et a l'argent du joueur
func (p *Personnage) Creation_Objet(nb, nbr int, str1, str2 string) {

	if p.verif_materiaux(nb, str1) && p.monnaie > nbr && p.CheckInventory() {
		for i := 1; i < nb; i++ {
			p.removeInventory(str1)
		}
		p.addInventory(str2)
		fmt.Println("vous avez frabriqué", str2)
		p.retrait_monnaie(nbr)

	} else if !p.verif_materiaux(nb, str1) {
		fmt.Println("vous n'avez pas assez de marteriaux pour faire cet objet")
	} else if p.monnaie < nbr {
		fmt.Println("vous n'avez pas assez de rubis pour payer la fabrication de l'objet")
	} else if !p.CheckInventory() {
		fmt.Println("vous n'avez pas assez de place dans votre inventaire")
	}

}

//Fonction qui sert de menu pour les equipement pour les fabriquer
func (p *Personnage) Menu_Equipement_Creation() {
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                Equipement          |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	fmt.Println("|             1.Fabriquer un         |")
	fmt.Println("|               casque               |")
	fmt.Println("|                                    |")
	fmt.Println("|             2. Fabriquer un        |")
	fmt.Println("|               plastron             |")
	fmt.Println("|                                    |")
	fmt.Println("|             3.Fabriquer une        |")
	fmt.Println("|               paire de             |")
	fmt.Println("|               bottes               |")
	fmt.Println("|                                    |")
	fmt.Println("|            4.Retour au menu        |")
	fmt.Println("|              précédent             |")
	fmt.Println("|____________________________________|")
	switch p.Scan() {
	case "1":
		p.Menu_casque_Creation()
	case "2":
		p.Menu_plastron_Creation()
	case "3":
		p.Menu_bottes_Creation()
	case "4":
		p.Menu_Fogeron()
	default:
		fmt.Println("Vous avez rentré une valeur éroné")
		p.Menu_Equipement_Creation()
	}
}

//Fonction qui sert de menu pour les casques pour les fabriquer
func (p *Personnage) Menu_casque_Creation() {
	var e1 Equipement
	e1.Init_List()
	i := 0
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                Casque              |")
	fmt.Println("|                Création            |")
	fmt.Println("|                                    |")
	for ; i < len(e1.liste_casque); i++ {
		fmt.Println("|             ", i+1, ".", e1.liste_casque[i], "   |")
		fmt.Println("|                                    |")
	}
	fmt.Println("|        ", i+1, "Revenir au ménu précédent            |")
	fmt.Println("|____________________________________|")
	switch p.Scan(){
	case "1":
		p.Creation_Objet(4, 15, "Cuir", e1.liste_casque[0])
		p.Menu_casque_Creation()
	case "2":
		p.Creation_Objet(4, 25, "Fer", e1.liste_casque[1])
		p.Menu_casque_Creation()
	case "3":
		p.Creation_Objet(4, 35, "Mitrhil", e1.liste_casque[2])
		p.Menu_casque_Creation()
	case "4":
		p.Menu_Equipement_Creation()
	default:
		fmt.Println("Vous avez rentré une valeur éroné")
		p.Menu_casque_Creation()
	}
}

//Fonction qui sert de menu pour les plastrons pour les fabriquer
func (p *Personnage) Menu_plastron_Creation() {
	var e1 Equipement
	e1.Init_List()
	i := 0
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                plastron            |")
	fmt.Println("|                Création            |")
	fmt.Println("|                8 Materiaux          ")
	fmt.Println("|                                    |")
	for ; i < len(e1.liste_plastron); i++ {
		fmt.Println("|             ", i+1, ".", e1.liste_plastron[i], "   |")
		fmt.Println("|                                    |")
	}
	fmt.Println("|        ", i+1, "Revenir au ménu précédent            |")
	fmt.Println("|____________________________________|")
	switch p.Scan() {
	case "1":
		p.Creation_Objet(8, 25, "Cuir", e1.liste_plastron[0])
		p.Menu_plastron_Creation()
	case "2":
		p.Creation_Objet(8, 35, "Fer", e1.liste_plastron[1])
		p.Menu_plastron_Creation()
	case "3":
		p.Creation_Objet(8, 45, "Mitrhil", e1.liste_plastron[2])
		p.Menu_plastron_Creation()
	case "4":
		p.Menu_Equipement_Creation()
	default:
		fmt.Println("Vous avez rentré une valeur éroné")
		p.Menu_plastron_Creation()
	}
}

//Fonction qui sert de menu pour les bottes pour les fabriquer
func (p *Personnage) Menu_bottes_Creation() {
	var e1 Equipement
	e1.Init_List()
	i := 0
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                bottes              |")
	fmt.Println("|                Création            |")
	fmt.Println("|                                    |")
	for ; i < len(e1.liste_bottes); i++ {
		fmt.Println("|             ", i+1, ".", e1.liste_bottes[i], "   |")
		fmt.Println("|                                    |")
	}
	fmt.Println("|        ", i+1, "Revenir au ménu précédent            |")
	fmt.Println("|____________________________________|")
	switch p.Scan() {
	case "1":
		p.Creation_Objet(4, 20, "Cuir", e1.liste_bottes[0])
		p.Menu_bottes_Creation()
	case "2":
		p.Creation_Objet(4, 30, "fer", e1.liste_bottes[1])
		p.Menu_bottes_Creation()
	case "3":
		p.Creation_Objet(4, 40, "mitrhil", e1.liste_bottes[2])
		p.Menu_bottes_Creation()
	case "4":
		p.Menu_Equipement_Creation()
	default:
		fmt.Println("Vous avez rentré une valeur éroné")
		p.Menu_bottes_Creation()

	}

}

//Fonction qui sert de menu pour les armes de mage pour les fabriquer
func (p *Personnage) Menu_arme_mage_Creation() {
	var e1 Equipement
	e1.Init_List()
	i := 0
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                Armes               |")
	fmt.Println("|                mage                |")
	fmt.Println("|                Création            |")
	fmt.Println("|                                    |")
	for ; i < len(e1.liste_armes_mage); i++ {
		fmt.Println("|             ", i+1, ".", e1.liste_armes_mage[i], "   |")
		fmt.Println("|                                    |")
	}
	fmt.Println("|        ", i+1, "Revenir au ménu précédent            |")
	fmt.Println("|____________________________________|")
	switch p.Scan() {
	case "1":
		p.Creation_Objet(4, 20, "Bois", e1.liste_armes_mage[0])
		p.Menu_arme_mage_Creation()
	case "2":
		p.Creation_Objet(4, 30, "Fer", e1.liste_armes_mage[1])
		p.Menu_arme_mage_Creation()
	case "3":
		p.Creation_Objet(4, 40, "Acier", e1.liste_armes_mage[2])
	case "4":
		p.Creation_Objet(4, 50, "Mitrhil", e1.liste_armes_mage[3])
		p.Menu_arme_mage_Creation()
	case "5":
		p.Menu_Equipement_Creation()
	default:
		fmt.Println("Vous avez rentré une valeur éroné")
		p.Menu_arme_mage_Creation()
	}
}

//Fonction qui sert de menu pour les armes d archer pour les fabriquer
func (p *Personnage) Menu_arme_archer_Creation() {
	var e1 Equipement
	i := 0
	e1.Init_List()
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                Armes               |")
	fmt.Println("|                archer              |")
	fmt.Println("|                Création            |")
	fmt.Println("|                                    |")
	for ; i < len(e1.liste_armes_archer); i++ {
		fmt.Println("|             ", i+1, ".", e1.liste_armes_archer[i], "   |")
		fmt.Println("|                                    |")
	}
	fmt.Println("|        ", i+1, "Revenir au ménu précédent            |")
	fmt.Println("|____________________________________|")
	switch p.Scan() {
	case "1":
		p.Creation_Objet(4, 20, "Bois", e1.liste_armes_archer[0])
		p.Menu_arme_archer_Creation()
	case "2":
		p.Creation_Objet(4, 30, "Fer", e1.liste_armes_archer[1])
		p.Menu_arme_archer_Creation()
	case "3":
		p.Creation_Objet(4, 40, "Acier", e1.liste_armes_archer[2])
	case "4":
		p.Creation_Objet(4, 50, "Mitrhil", e1.liste_armes_archer[3])
		p.Menu_arme_archer_Creation()
	case "5":
		p.Menu_arme_archer_Creation()
	default:
		fmt.Println("Vous avez rentré une valeur éroné")
		p.Menu_arme_archer_Creation()
	}
}

//Fonction qui sert de menu pour les armes d'épéiste pour les fabriquer
func (p *Personnage) Menu_arme_épéiste_Creation() {
	var e1 Equipement
	e1.Init_List()
	i := 0
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                Armes               |")
	fmt.Println("|                épéiste             |")
	fmt.Println("|                Création            |")
	fmt.Println("|                                    |")
	for ; i < len(e1.liste_armes_épéiste); i++ {
		fmt.Println("|             ", i+1, ".", e1.liste_armes_épéiste[i], "   |")
		fmt.Println("|                                    |")
	}
	fmt.Println("|        ", i+1, "Revenir au ménu précédent            |")
	fmt.Println("|____________________________________|")
	switch p.Scan(){
	case "1":
		p.Creation_Objet(4, 20, "Bois", e1.liste_armes_épéiste[0])
		p.Menu_arme_épéiste_Creation()
	case "2":
		p.Creation_Objet(4, 30, "Fer", e1.liste_armes_épéiste[1])
		p.Menu_arme_épéiste_Creation()
	case "3":
		p.Creation_Objet(4, 40, "Acier", e1.liste_armes_épéiste[2])
		p.Menu_arme_épéiste_Creation()
	case "4":
		p.Creation_Objet(4, 50, "Mitrhil", e1.liste_armes_épéiste[3])
		p.Menu_arme_épéiste_Creation()
	case "5":
		p.Menu_Equipement_Creation()
	default:
		fmt.Println("Vous avez rentré une valeur éroné")
		p.Menu_arme_épéiste_Creation()
	}
}
