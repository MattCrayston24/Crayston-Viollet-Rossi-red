package Dossier_Package

import "fmt"

// Fonction pour définir le menu :
func (p Personnage) Menu() {
	var choix_menu string
	for {
		fmt.Print("Taper le numéro du menu dans lequel vous voulez rentrée: \n1 :inventaire \n2 :statistique \n3 :forge\n4 :Quitter:\n")
		fmt.Scan(&choix_menu)

		switch {
		case choix_menu == "1":
			p.AccèsInventaire()
		case choix_menu == "2":
			p.Afficher_info()
		case choix_menu == "3":
			p.Menu_Fogeron()
		case choix_menu == "4":

		default:
			fmt.Println("Vous n'avez pas selectionner un reponse valide")
		}
	}
}

func (p Personnage) Menu_Fogeron() {
	var objet_choisi string
	fmt.Println("Taper le nom de l'objet que vous voulez fabriquer : \nChapeau en cuir \nPlastron en cuir\nJambière en cuir ")
	fmt.Scan(&objet_choisi)

	switch {
	case objet_choisi == "Chapeau en cuir":
		p.Creation_Objet(2, 5, "cuir", "Chapeau en cuir")

	case objet_choisi == "plastron en cuir":
		p.Creation_Objet(3, 10, "cuir", "Plastron en cuir")

	case objet_choisi == "Jambière en cuir":
		p.Creation_Objet(2, 5, "cuir", "Jambière en cuir")

	default:
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
