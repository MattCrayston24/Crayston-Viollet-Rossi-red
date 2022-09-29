package Dossier_Package

import (
	"bufio"
	"fmt"
	"os"
)

// Structure du personnage avec ses attributs et ses méthodes
type Personnage struct {
	nom                  string
	classe               string
	point_de_vie_maximum int
	point_de_vie_actuel  int
	skill                Skill
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
	attaque_base         int
	point_skill          int
}

// Fonction init pour créer un personnage avec son pseudo
func (p *Personnage) Init() {
	for {
		fmt.Println("______________________________________________________________________________")
		fmt.Println("| Choisissez votre nom comprenant que des lettres (pas d'espace ni d'accent):|")
		fmt.Println("|____________________________________________________________________________|")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			line := scanner.Text()
			p.nom = line
		}
		if p.verif_nom(p.nom) {
			p.nom = p.majuscule(p.nom)
			fmt.Println("Votre nom est ", p.nom)
			break
		}
		fmt.Println("Votre nom est invalide")
	}

	p.choix_classe()

}

//Fonction qui affiche les statistiques du personnage
func (p Personnage) Afficher_info() {
	fmt.Println("______________________________________")
	fmt.Println("|                                     ")
	fmt.Println("|              Statistique            ")
	fmt.Println("|_____________________________________")
	fmt.Println("|                                     ")
	fmt.Println("|     pseudo : ", p.nom, "              ")
	fmt.Println("|                                     ")
	fmt.Println("|     vie : ", p.point_de_vie_actuel, "/", p.point_de_vie_maximum, "                         ")
	fmt.Println("|                                     ")
	fmt.Println("|     mana : ", p.mana_actuel, "/", p.mana_maximum, "                         ")
	fmt.Println("|                                     ")
	fmt.Println("|     skill : ", p.skill.nom, "                         ")
	fmt.Println("|                                     ")
	fmt.Println("|       Monnaie :", p.monnaie, "                   ")
	fmt.Println("|                                     ")
	fmt.Println("|       Classe :", p.classe, "                   ")
	fmt.Println("|                                     ")
	fmt.Println("|        Attaque :", p.attaque_base, "                             ")
	fmt.Println("|                                     ")
	fmt.Println("|        Experience :", p.experience_actuel, "/", p.experience_max, "                             ")
	fmt.Println("|_____________________________________")
	p.Menu()
}

//Fonction qui retire de la monnaie au personnage 
func (p *Personnage) retrait_monnaie(nb int) {
	p.monnaie -= nb

}

//Fonction qui ajoute de la monnaie au personnage
func (p *Personnage) ajout_monnaie(nb int) {
	p.monnaie += nb
}

//Fonction qui vérifie si les points de vie du personnage sont inférieur ou égal à 0
func (p *Personnage) wasted() bool {
	return p.point_de_vie_actuel <= 0
}

//Fonction qui permet a l'utilisateur de choisir sa classe et envoie les statistiques du personnage en fonction de la classe choisie
func (p *Personnage) choix_classe() {
	var classe_choisi string
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
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		classe_choisi = line
	}
	for i := 0; i < 1; i++ {
		switch classe_choisi {
		case "1":
			p.Archer()
			fmt.Println("Votre classe est : ", p.classe)
			break
		case "2":
			p.Epéiste()
			fmt.Println("Votre classe est : ", p.classe)
			break
		case "3":
			p.Mage()
			fmt.Println("Votre classe est : ", p.classe)
			break
		default:
			fmt.Println("Erreur le caractère entré n'est pas valide")
			p.choix_classe()
		}

		fmt.Println("______________________________________")
		fmt.Println("|                                     ")
		fmt.Println("|              Statistique            ")
		fmt.Println("|_____________________________________")
		fmt.Println("|                                     ")
		fmt.Println("|     pseudo : ", p.nom, "              ")
		fmt.Println("|                                     ")
		fmt.Println("|     vie : ", p.point_de_vie_actuel, "/", p.point_de_vie_maximum, "                         ")
		fmt.Println("|                                     ")
		fmt.Println("|     mana : ", p.mana_actuel, "/", p.mana_maximum, "                         ")
		fmt.Println("|                                     ")
		fmt.Println("|     skill : ", p.skill.nom, "                         ")
		fmt.Println("|                                     ")
		fmt.Println("|       Monnaie :", p.monnaie, "                   ")
		fmt.Println("|                                     ")
		fmt.Println("|       Classe :", p.classe, "                   ")
		fmt.Println("|_____________________________________")

	}
}
