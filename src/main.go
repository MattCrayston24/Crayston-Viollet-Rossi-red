package main

import (
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> 615e65fd52588bdaf6b94ae0fc2483c57b54862a
	"src/Dossier_Package"
)

func main() {
	var p1 Dossier_Package.Personnage
	p1.Init()

	Dossier_Package.Hello()
}
<<<<<<< HEAD
=======

// Fonction pour définir le menu :
func Menu() {
	var nombre_choisi string
	for {
		fmt.Print("Si vous voulez accédez a l'inventaire taper 1, \n Si vous voulez accédez aux statistiques de votre personnage taper 2, \n Si vous voulez quitter taper 3 : ")
		fmt.Scan(&nombre_choisi)
		if nombre_choisi == "1" {
			fmt.Println("salut")
			break
		} else if nombre_choisi == "2" {
			fmt.Println("coucou")
			break
		} else if nombre_choisi == "3" {
			break
		}
	}
}
>>>>>>> 615e65fd52588bdaf6b94ae0fc2483c57b54862a
