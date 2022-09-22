package Dossier_Package

import (
	"fmt"
	"time"
)

type Equipement struct {
	casque   string
	plastron string
	bottes   string
}

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
}

// Fonction init pour créer un personnage :
func (p *Personnage) Init() {
	var classe_choisi string
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
	for {
		fmt.Print("Choisissez votre classe, \n Si vous voulez la classe tank taper 1, \n Si vous voulez la classe attaquant taper 2, \n Si vous voulez la calsse équilibré tapez 3 :")
		fmt.Scan(&classe_choisi)
		if classe_choisi == "1" {
			p.classe = "tank"
			fmt.Print("Votre classe est : ", p.classe)
			break
		} else if classe_choisi == "2" {
			p.classe = "attaquant"
			fmt.Print("Votre classe est : ", p.classe)
			break
		} else if classe_choisi == "3" {
			p.classe = "equilibré"
			fmt.Print("Votre classe est : ", p.classe)
			break
		}
	}
	if p.classe == "tank" {
		p.Tank()
	} else if p.classe == "attaquant" {
		p.Attaquant()
		p.Skill = append(p.Skill, "coup de épée")
	} else if p.classe == "equilibré" {
		p.Equilibré()
	}
	p.taille_inventaire = 10
	p.addInventory("potion")
	p.niveau = 1
	p.monnaie = 100
	fmt.Println("\n Vos point de vie maximum sont : ", p.point_de_vie_maximum)
	fmt.Println("Vos points de vie actuel sont : ", p.point_de_vie_actuel)
	fmt.Println("votre skil est :", p.Skill)
	fmt.Println("Votre niveau actuel est : ", p.niveau)
	fmt.Println("Vous avez  : ", p.monnaie, " rubis")
}

func (p *Personnage) TakePot() {
	potion := 20
	for i := range p.inventaire {
		if p.inventaire[i] == "potion" {
			if p.point_de_vie_actuel+potion > p.point_de_vie_maximum {
				p.point_de_vie_actuel = p.point_de_vie_maximum
				p.inventaire = append(p.inventaire[:i], p.inventaire[i+1:]...)
				break
			} else {
				p.point_de_vie_actuel += potion
				p.inventaire = append(p.inventaire[:i], p.inventaire[i+1:]...)
				break
			}
		}
	}
}

func (p *Personnage) PoisonPot() {
	p.point_de_vie_actuel -= 10
	time.Sleep(1 * time.Second)
	p.point_de_vie_actuel -= 10
	time.Sleep(1 * time.Second)
	p.point_de_vie_actuel -= 10
}

func (p Personnage) Afficher_info() {
	fmt.Println("Votre nom est :", p.nom)
	fmt.Println("Votre classe est:", p.classe)
	fmt.Println("Vos point de vie max sont :", p.point_de_vie_maximum)
	fmt.Println("Vos point de vie actuel sont :", p.point_de_vie_actuel)
	fmt.Println("vos skill sont :", p.Skill)
}

func (p *Personnage) retrait_monnaie(nb int) {
	p.monnaie -= nb

}

func (p *Personnage) ajout_monnaie(nb int) {
	p.monnaie += nb
}
