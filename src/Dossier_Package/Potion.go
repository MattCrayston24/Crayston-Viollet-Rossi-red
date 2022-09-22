package Dossier_Package

import (
	"fmt"
	"time"
)

func (p *Personnage) TakePot() {
	fmt.Println("vous aviez", p.point_de_vie_actuel, "/", p.point_de_vie_maximum)
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
	fmt.Println("Vous avez maintenant", p.point_de_vie_actuel, "/", p.point_de_vie_maximum)
}

func (p *Personnage) PoisonPot() {
	p.point_de_vie_actuel -= 10
	time.Sleep(1 * time.Second)
	p.point_de_vie_actuel -= 10
	time.Sleep(1 * time.Second)
	p.point_de_vie_actuel -= 10
}
