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
	fmt.Println("|             4. Cuir                |")
	fmt.Println("|               (15 monnaie)         |")
	fmt.Println("|                                    |")
	fmt.Println("|             5.Fer                  |")
	fmt.Println("|               (25 monnaie)         |")
	fmt.Println("|                                    |")
	fmt.Println("|             6. Metal               |")
	fmt.Println("|               (35 monnaie)         |")
	fmt.Println("|                                    |")
	fmt.Println("|             7. Mitrhil             |")
	fmt.Println("|               (45 monnaie)         |")
	fmt.Println("|                                    |")
	fmt.Println("|             8.Quitter              |")
	fmt.Println("|____________________________________|")
	fmt.Scan(&item_choisi)
	switch item_choisi {
	case 1:
		p.addInventory("Potion de Poison")
		p.retrait_monnaie(30)
		fmt.Println("La potion a été ajouté a votre invertaire")
		p.Marchand()
	case 2:
		p.addInventory("Potion de vie")
		p.retrait_monnaie(20)
		fmt.Println("La potion a été ajouté a votre invertaire")
		p.Marchand()
	case 3:
		p.taille_inventaire += 3
		p.retrait_monnaie(500)
		fmt.Println("Les places supplémentaires ont été ajoutées")
		p.Marchand()
	case 4:
		p.addInventory("Cuir")
		p.retrait_monnaie(15)
		fmt.Println("Le Cuir a été ajouter a votre inventaire")
		p.Marchand()
	case 5:
		p.addInventory("Fer")
		p.retrait_monnaie(25)
		fmt.Println("Le Fer a été ajouter a votre inventaire")
		p.Marchand()
	case 6:
		p.addInventory("Metal")
		p.retrait_monnaie(35)
		fmt.Println("Le Metal a été ajouter a votre inventaire")
		p.Marchand()
	case 7:
		p.addInventory("Mithril")
		p.retrait_monnaie(45)
		fmt.Println("Le Mithril a été ajouter a votre inventaire")
		p.Marchand()
	case 8:
		fmt.Println("Vous quiitez le marchand")
		p.Menu()
	default:
		fmt.Println("Vous n'avez pas ecrit un reponse valide ")
		p.Marchand()
	}
}
