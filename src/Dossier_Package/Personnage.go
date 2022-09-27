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
		fmt.Print("Choisissez votre nom comprenant au moins 3 caractères :")
		fmt.Scan(&nom_choisi)
		if len(nom_choisi) >= 3 {
			fmt.Print("Votre nom est : ", nom_choisi)
			p.nom = nom_choisi
			break
		}
	}
	fmt.Print("Choisissez votre classe, \n Si vous voulez la classe Archer taper 1, \n Si vous voulez la classe épéiste taper 2, \n Si vous voulez la calsse mage tapez 3 :")
	fmt.Scan(&classe_choisi)
	for i := 0; i < 1; i++ {
		switch classe_choisi {
		case 1:
			p.Archer()
			fmt.Print("Votre classe est : ", p.classe)
		case 2:
			p.Epéiste()
			fmt.Print("Votre classe est : ", p.classe)
		case 3:
			p.Mage()
			fmt.Print("Votre classe est : ", p.classe)
		default:
			fmt.Println("taper un reponse valide ")
			p.Init()
		}
	}
	fmt.Println("\n Vos point de vie maximum sont : ", p.point_de_vie_maximum)
	fmt.Println("Vos points de vie actuel sont : ", p.point_de_vie_actuel)
	fmt.Println("Votre réserve de mana est :", p.mana_actuel, "/", p.mana_maximum)
	fmt.Println("vos skil sont :", p.Skill)
	fmt.Println("Votre niveau actuel est : ", p.niveau)
	fmt.Println("Vous avez  : ", p.monnaie, " rubis")
}

func (p Personnage) Afficher_info() {
	fmt.Println("Votre nom est :", p.nom)
	fmt.Println("Votre classe est:", p.classe)
	fmt.Println("Vos point de vie max sont :", p.point_de_vie_maximum)
	fmt.Println("Vos point de vie actuel sont :", p.point_de_vie_actuel)
	fmt.Println("vos skill sont :", p.Skill)
	p.Menu()
}

func (p *Personnage) retrait_monnaie(nb int) {
	p.monnaie -= nb

}

func (p *Personnage) ajout_monnaie(nb int) {
	p.monnaie += nb
}
