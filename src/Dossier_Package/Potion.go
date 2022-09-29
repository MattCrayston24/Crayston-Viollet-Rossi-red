package Dossier_Package

import (
	"fmt"
	"time"
)

func (p *Personnage) TakePot() {
	fmt.Println("vous aviez", p.point_de_vie_actuel, "/", p.point_de_vie_maximum)
	potion := 20
	for i := range p.inventaire {
		if p.inventaire[i] == "Potion de Vie" {
			if p.point_de_vie_actuel+potion > p.point_de_vie_maximum {
				p.point_de_vie_actuel = p.point_de_vie_maximum
				p.removeInventory("Potion de Vie")
				break
			} else {
				p.point_de_vie_actuel += potion
				p.removeInventory("Potion de Vie")
				break
			}
		}
	}
	fmt.Println("Vous avez maintenant", p.point_de_vie_actuel, "/", p.point_de_vie_maximum)
}

func (p *Personnage) PoisonPot(m *Monstre) {
	for i := range p.inventaire {
		if p.inventaire[i] == "Potion de poison" {
			m.point_de_vie_actuel -= 10

			time.Sleep(1 * time.Second)
			m.point_de_vie_actuel -= 10
			time.Sleep(1 * time.Second)
			m.point_de_vie_actuel -= 10
			fmt.Print("Les nouveau points de vie du gobelin sont :", m.point_de_vie_actuel, "/", m.point_de_vie_max)
			break
		}
	}
}
