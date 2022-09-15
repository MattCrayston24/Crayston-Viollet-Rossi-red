package main

import "fmt"

func main(){
	var p1 Personnage
	p1.Init()
}


// Fonction pour définir le menu :
func Menu(){
	var nombre_choisi string
	for{	
		fmt.Print("Si vous voulez accédez a l'inventaire taper 1, \n Si vous voulez accédez aux statistiques de votre personnage taper 2, \n Si vous voulez quitter taper 3 : ")
		fmt.Scan(&nombre_choisi)
		if nombre_choisi == "1" {
			fmt.Println("salut")
			break
		}else if nombre_choisi == "2"{
			fmt.Println("coucou")
			break
		}else if nombre_choisi == "3"{
			break
		}
	}
}


// Définition d'une structure :
type Personnage struct{
	nom string
	classe string
	point_de_vie_maximum int 
	point_de_vie_actuel int
	inventaire []string
}


// Fonction init pour créer un personnage :
func (p *Personnage) Init(){ 
	var classe_choisi string
	var nom_choisi string
	for{
		fmt.Print("Choisissez votre nom comprenant au moins 3 caractères :")
		fmt.Scan(&nom_choisi)
		if len(nom_choisi) >= 3 {
			fmt.Print("Votre nom est : ", nom_choisi)
			p.nom = nom_choisi
			break
		}
	}
	for{
		fmt.Print("Choisissez votre classe, \n Si vous voulez la classe tank taper 1, \n Si vous voulez la classe attaquant taper 2, \n Si vous voulez la calsse équilibré tapez 3 :")
		fmt.Scan(&classe_choisi)
		if classe_choisi == "1"{
			p.classe = "tank"
			fmt.Print("Votre classe est : ",p.classe)
			break
		}else if classe_choisi == "2"{
			p.classe = "attaquant"
			fmt.Print("Votre classe est : ",p.classe)
			break
		}else if classe_choisi == "3"{
			p.classe = "equilibré"
			fmt.Print("Votre classe est : ",p.classe)
			break
		}
	}
		if p.classe == "tank"{
		p.point_de_vie_maximum = 100
		p.point_de_vie_actuel = 100
	}else if p.classe == "attaquant"{ 
		p.point_de_vie_maximum = 50
		p.point_de_vie_actuel = 50
	}else if p.classe == "equilibré"{ 
		p.point_de_vie_maximum = 75
		p.point_de_vie_actuel = 75
	}
	fmt.Println("\n Vos point de vie maximum sont : ",p.point_de_vie_maximum)
	fmt.Println(" Vos points de vie actuel sont : ",p.point_de_vie_actuel)
}


