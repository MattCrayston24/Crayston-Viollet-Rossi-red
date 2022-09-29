package Dossier_Package

import "fmt"

func (p *Personnage) addExp(nb int) {
	p.experience_actuel += nb
	if p.experience_actuel > p.experience_max {
		p.experience_actuel = p.experience_actuel - p.experience_max
		p.addLevev()
		fmt.Println("vous avez montez de  niveaux vous êtes maintenant niveau", p.niveau)
	}
}

func (p *Personnage) addLevev() {
	p.experience_max += 100
	p.niveau += 1
	p.point_de_vie_maximum += 25
	p.attaque_base += 2
	p.point_skill += 1

}
