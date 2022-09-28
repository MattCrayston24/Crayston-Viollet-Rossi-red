package Dossier_Package

import (
	"fmt"
)

// Définition d'une structure :
type Personnage struct {
	nom                  string
	classe               string
	point_de_vie_maximum int
	point_de_vie_actuel  int
	Skill                []string
	inventaire           []string
	niveau               int
	monnaie              int
	mon_equipement       Equipement
	taille_inventaire    int
	points_attaque       int
	initiative           int
	experience_actuel    int
	experience_max       int
	mana_maximum         int
	mana_actuel          int
}

// Fonction init pour créer un personnage :
func (p *Personnage) Init() {
	var classe_choisi int
	var nom_choisi string
	for {
		fmt.Println("_____________________________________________________________________________")
		fmt.Println("| Choisissez votre nom comprenant au moins 3 caractères et que des lettres :|")
		fmt.Println("|___________________________________________________________________________|")
		fmt.Scan(&nom_choisi)
		if p.verif_nom(nom_choisi) == true {
			p.nom = p.majuscule(nom_choisi)
			fmt.Println("Votre nom est :", p.nom)
			break
		}	
		fmt.Println("Votre nom est invalide")
	}
	for i := 0 ; i<1 ; i++  {	
		fmt.Println("______________________________________")
		fmt.Println("|                                    |")
		fmt.Println("|                Choix               |")
		fmt.Println("|                Classe              |")
		fmt.Println("|____________________________________|")
		fmt.Println("|                                    |")
		fmt.Println("|             1.Archer               |")
		fmt.Println("|                                    |")
		fmt.Println("|             2.Epéiste              |")
		fmt.Println("|                                    |")
		fmt.Println("|             3.Mage                 |")
		fmt.Println("|____________________________________|")
		fmt.Scan(&classe_choisi)
		switch classe_choisi {
		case 1:
			p.Archer()
			fmt.Println("Votre classe est : ", p.classe)
		case 2:
			p.Epéiste()
			fmt.Println("Votre classe est : ", p.classe)
		case 3:
			p.Mage()
			fmt.Println("Votre classe est : ", p.classe)
		default:
			fmt.Println("Erreur le caractère entré n'est pas valide")
			i--
		}	
					
	}
	fmt.Println("______________________________________")
	fmt.Println("|                                     ")
	fmt.Println("|              Statistique            ")
	fmt.Println("|_____________________________________")
	fmt.Println("|                                     ")
	fmt.Println("|     pseudo : ",p.nom,"              ")
	fmt.Println("|                                     ")
	fmt.Println("|     vie : ",p.point_de_vie_actuel,"/",p.point_de_vie_maximum ,"                         ")
	fmt.Println("|                                     ")
	fmt.Println("|     mana : ",p.mana_actuel,"/",p.mana_maximum ,"                         ")
	fmt.Println("|                                     ")
	fmt.Println("|     skill : ",p.Skill,"                         ")                                                              
	fmt.Println("|_____________________________________")	
}

func (p Personnage) Afficher_info() {
	fmt.Println("______________________________________")
	fmt.Println("|                                     ")
	fmt.Println("|              Statistique            ")
	fmt.Println("|_____________________________________")
	fmt.Println("|                                     ")
	fmt.Println("|     pseudo : ",p.nom,"              ")
	fmt.Println("|                                     ")
	fmt.Println("|     vie : ",p.point_de_vie_actuel,"/",p.point_de_vie_maximum ,"")
	fmt.Println("|                                     ")
	fmt.Println("|     mana : ",p.mana_actuel,"/",p.mana_maximum ,"")
	fmt.Println("|                                     ")
	fmt.Println("|     skill : ",p.Skill,"             ")                                                              
	fmt.Println("|_____________________________________")
	p.Menu()
}

func (p *Personnage) retrait_monnaie(nb int) {
	p.monnaie -= nb

}

func (p *Personnage) ajout_monnaie(nb int) {
	p.monnaie += nb
}