package Dossier_Package

import "fmt"

// Fonction pour définir le menu :
func (p Personnage) Menu() {
	var choix_menu string
	for {
		fmt.Print("Taper le numéro du menu dans lequel vous voulez rentrée: \n1 :inventaire \n2 :statistique \n3 :forge\n4 : Marchand \n5:Quitter:\n")
		fmt.Scan(&choix_menu)

		switch {
		case choix_menu == "1":
			p.AccèsInventaire()
		case choix_menu == "2":
			p.Afficher_info()
		case choix_menu == "3":
			p.Menu_Fogeron()
		case choix_menu == "4":
			p.Marchand()
		case choix_menu == "5":
			break
		default:
			fmt.Println("Vous n'avez pas selectionner un reponse valide")
		}
	}
}

func (p Personnage) Menu_Fogeron() {
	var objet_choisi string
	fmt.Println("Taper le numero de l'objet que vous voulez fabriquer : \n1 :Chapeau en cuir \n2 :Plastron en cuir\n3: Jambière en cuir ")
	fmt.Scan(&objet_choisi)
	for {
		switch {
		case objet_choisi == "1":
			p.Creation_Objet(2, 5, "cuir", "Chapeau en cuir")

		case objet_choisi == "2":
			p.Creation_Objet(3, 10, "cuir", "Plastron en cuir")

		case objet_choisi == "3":
			p.Creation_Objet(2, 5, "cuir", "Jambière en cuir")

		default:
			break
		}
	}
}

func (p Personnage) Creation_Objet(nb, nbr int, str1, str2 string) {

	if p.verif_materiaux(nb, str1) && p.monnaie > nbr && p.CheckInventory() {
		for i := 0; i < nb; i++ {
			p.removeInventory(str1)
		}
		p.addInventory(str2)
		p.retrait_monnaie(nbr)

	} else if !p.verif_materiaux(nb, str1) {
		fmt.Println("vous n'avez pas assez de marteriaux pour faire cet objet")
	} else if p.monnaie < nbr {
		fmt.Println("vous n'avez pas assez de rubis pour payer la fabrication de l'objet")
	} else if !p.CheckInventory() {
		fmt.Println("vous n'avez pas assez de place dans votre inventaire")
	}

}

func (p Personnage) Menu_Inventaire() {
	var choix string
	fmt.Println("taper le numéro de l'action que vous voulez faire : \n1 : Prendre potion de vie \n2 :Mettre un équipement")
	fmt.Scan(&choix)
	for {
		switch {
		case choix == "1":
			p.TakePot()
			p.Menu()
		case choix == "2":
			fmt.Println("taper le numéro correspondant pour équiper l'objet choisi : \n1 :casque en cuir :\n2: plastron de cuir \n3: jambière en cuir")
			p.mettre_equipement(choix)
			p.Ajout_Stat_equipement()
		default:

		}
	}
}
