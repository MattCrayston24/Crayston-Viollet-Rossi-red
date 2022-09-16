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

/*
if nombre_choisi == "1" {
	fmt.Println("salut")
	break
} else if nombre_choisi == "2" {
	fmt.Println("coucou")
	break
} else if nombre_choisi == "3" {
	break
}*/
