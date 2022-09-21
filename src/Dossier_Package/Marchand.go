package Dossier_Package

import "fmt"

func (p *Personnage) Marchand() {
	var item_choisi string
	fmt.Println(" Taper le numéro du Produit que vous voulez acheter :\n 1:Potion de Poison \n2: Potion de vie \n3: Acheter slot invetaire")
	fmt.Scan(&item_choisi)
	for {
		switch {
		case item_choisi == "1":
			p.addInventory("Potion de Poison")
			p.retrait_monnaie(30)
		case item_choisi == "2":
			p.addInventory("Potion de vie")
			p.retrait_monnaie(20)
			fmt.Println("La potion a été ajouté a votre invertaire")
		case item_choisi == "3":
			p.retrait_monnaie(500)
		default:
			fmt.Println("Vous n'avez pas ecrit un reponse valide ")
		}
	}
}
