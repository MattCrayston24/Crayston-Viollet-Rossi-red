package Dossier_Package

import "fmt"

// Fonction pour définir le menu :
func (p Personnage) Menu() {
	var choix_menu int
	fmt.Print("Taper le numéro du menu dans lequel vous voulez entrée: \n1 :inventaire \n2 :statistique \n3 :forge\n4 :Marchand \n5: Combat entrainement \n6:Quitter\n")
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
		break
	default:
		fmt.Println("Vous n'avez pas selectionner un reponse valide")
	}
}

func (p Personnage) Menu_Fogeron() {
	var objet_choisi int
	fmt.Println("Taper le numero de l'objet que vous voulez fabriquer : \n1 :Chapeau en cuir \n2 :Plastron en cuir\n3: Jambière en cuir ")
	fmt.Scan(&objet_choisi)
	switch objet_choisi {
	case 1:
		p.Creation_Objet(2, 5, "cuir", "Chapeau en cuir")

	case 2:
		p.Creation_Objet(3, 10, "cuir", "Plastron en cuir")

	case 3:
		p.Creation_Objet(2, 5, "cuir", "Jambière en cuir")

	default:
		break
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
	var choix int
	fmt.Println("taper le numéro de l'action que vous voulez faire : \n1 : Prendre potion de vie \n2 :Mettre un équipement \n3 : retour au menu précedent")
	fmt.Scan(&choix)
	switch choix {
	case 1:
		p.TakePot()
		p.Menu()
	case 2:
		p.Menu_Equipement()
		p.Ajout_Stat_equipement()
	case 3:
		p.Menu()
	default:

	}
}

func (p Personnage) Menu_Equipement() {
	var choix int
	fmt.Println("taper le numéro de l'action que vous voulez faire : \n1 : mettre un casque \n2 :Mettre un plastron \n3: mettre des bottes \n4 :retour au menu précedent")
	fmt.Scan(&choix)
	switch choix {
	case 1:
		p.Menu_casque()
	case 2:
		p.Menu_plastron()
	case 3:
		p.Menu_plastron()
	case 4:
		p.AccèsInventaire()
	}
}

func (p Personnage) Menu_casque() {
	var choix int
	fmt.Println("taper le numéro de l'action que vous voulez faire : \n1 : mettre un casque en cuir \n2: retour au menu précedent ")
	fmt.Scan(&choix)
	switch choix {
	case 1:
		p.Mettre_casque(1)
	case 2:
		p.Menu_Equipement()
	default:

	}

}

func (p Personnage) Menu_plastron() {
	var choix int
	fmt.Println("taper le numéro de l'action que vous voulez faire : \n1 : mettre un plastron en cuir \n2: retour au menu précedent")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		p.Mettre_Plastron(1)
	case 2:
		p.Menu_Equipement()
	default:

	}
}

func (p Personnage) Menu_bottes() {
	var choix int
	fmt.Println("taper le numéro de l'action que vous voulez faire : \n1 : mettre un bottes en cuir \n2: retour au menu précedent")
	fmt.Scan(&choix)
	switch choix {
	case 1:
		p.Mettre_bottes(1)
	case 2:
		p.Menu_Equipement()
	default:

	}
}
