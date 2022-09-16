package Dossier_Package

import "fmt"

// Fonction pour définir le menu :
func (p Personnage) Menu() {
	var nombre_choisi string
	for {
		fmt.Print("Si vous voulez accédez a l'inventaire taper 1, \n Si vous voulez accédez aux statistiques de votre personnage taper 2, \n Si vous voulez quitter taper 3 : ")
		fmt.Scan(&nombre_choisi)

		switch {
		case nombre_choisi == "1":
			p.AccèsInventaire()
			fmt.Println("salut")
		case nombre_choisi == "2":
			p.Afficher_info()
		case nombre_choisi == "3":

		default:
			fmt.Println("Vous n'avez pas selectionner un reponse valide")
		}
	}
}

func (p Personnage) Menu_Fogeron() {
	var objet_choisi string
	fmt.Println("Taper le nom de l'objet que vous voulez fabriquer : Chapeau en cuir \n Plastron en cuir\n jambière en cuir ")
	fmt.Scan(&objet_choisi)

	switch {
	case objet_choisi == "Chapeau en cuir":

	case objet_choisi == "plastron en cuir":

	case objet_choisi == "jambière en cuir":
	}
}
