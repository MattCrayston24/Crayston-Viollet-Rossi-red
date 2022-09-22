package Dossier_Package

import "fmt"

func (p *Personnage) Marchand() {
	var item_choisi string
	if len(p.inventaire) < 10 {
		fmt.Println("Potion de Poison taper 1 \n Potion de vie taper 2 \n Gain de trois places dans l'inventaire taper 3 ")
		fmt.Scan(&item_choisi)
		if item_choisi == "1" {
			p.addInventory("Potion de Poison")
			p.monnaie -= 30
		} else if item_choisi == "2" {
			p.addInventory("Potion de vie")
			p.monnaie -= 20
		} else if item_choisi == "3" {
			p.taille_inventaire += 3
		}
	}
}
