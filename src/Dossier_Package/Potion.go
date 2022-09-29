package Dossier_Package

import (
	"fmt"
	"time"
)

//Fonction qui permet de prendre une potion de vie si elle est dans l'inventaire et la supprimer de l'inventaire après utilisation 
func (p *Personnage) TakePot() {
	fmt.Println("vous aviez", p.point_de_vie_actuel, "/", p.point_de_vie_maximum)
	potion := 20
	for i := range p.inventaire {
		if p.inventaire[i] == "Potion de Vie" {
			if p.point_de_vie_actuel+potion > p.point_de_vie_maximum {
				p.point_de_vie_actuel = p.point_de_vie_maximum
				p.removeInventory("Potion de Vie")
				break
			} else if p.point_de_vie_actuel+potion <= p.point_de_vie_maximum {
				p.point_de_vie_actuel += potion
				p.removeInventory("Potion de Vie")
				break
			}else{
				fmt.Println("vous n'avez pas de potion de vie")
			}
		}
	}
	fmt.Println("Vous avez maintenant", p.point_de_vie_actuel, "/", p.point_de_vie_maximum)
}

//Fonction qui permet d'utiliser une potion de poison contre l'ennemi si elle est dans l'inventaire et la supprimer de l'inventaire après utilisation
func (p *Personnage) PoisonPot(m *Monstre) {
	count := 0
	for i := range p.inventaire {
		if p.inventaire[i] == "Potion de poison" {
			count++
			m.point_de_vie_actuel -= 10
			time.Sleep(1 * time.Second)
			m.point_de_vie_actuel -= 10
			time.Sleep(1 * time.Second)
			m.point_de_vie_actuel -= 10
			fmt.Print("Les nouveau points de vie du gobelin sont :", m.point_de_vie_actuel, "/", m.point_de_vie_max)
			break
		}
	}
	if count == 0 {
		fmt.Println("Vous n'avez pas de potion de poison")
	}
}
