package Dossier_Package

func (p *Personnage) addExp(nb int) {
	p.experience_actuel += nb
	if p.experience_actuel > p.experience_max {
		p.experience_actuel = p.experience_actuel - p.experience_max
		p.addLevev()
	}
}

func (p *Personnage) addLevev() {
	p.experience_max += 100
	p.niveau += 1
	p.point_de_vie_maximum += 25
	p.attaque_base += 2
	p.point_skill += 1
}