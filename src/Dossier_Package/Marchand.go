package Dossier_Package

import "fmt"

func (p *Personnage) Marchand() {
	var item_choisi int
	fmt.Println("_____________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|                Menu                |")
	fmt.Println("|                Marchand            |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	fmt.Println("|             1.Potion de Poison     |")
	fmt.Println("|               (30 monnaie)         |")
	fmt.Println("|                                    |")
	fmt.Println("|             2.Potion de vie        |")
	fmt.Println("|               (20 monnaie)         |")
	fmt.Println("|                                    |")
	fmt.Println("|             3.3 Places inventaire  |")
	fmt.Println("|               (500 monnaie)        |")
	fmt.Println("|                                    |")
	fmt.Println("|             4.Quitter              |")
	fmt.Println("|____________________________________|")
	fmt.Scan(&item_choisi)
	switch item_choisi {
	case 1:
		p.addInventory("Potion de Poison")
		p.retrait_monnaie(30)
		fmt.Println("La potion a été ajouté a votre invertaire")
		p.Menu()
	case 2:
		p.addInventory("Potion de vie")
		p.retrait_monnaie(20)
		fmt.Println("La potion a été ajouté a votre invertaire")
		p.Menu()
	case 3:
		p.taille_inventaire += 3
		p.retrait_monnaie(500)
		fmt.Println("Les places supplémentaires ont été ajoutées")
		p.Menu()
	case 4:
		fmt.Println("Vous quiitez le marchand")
		p.Menu()
	default:
		fmt.Println("Vous n'avez pas ecrit un reponse valide ")
	}
}
