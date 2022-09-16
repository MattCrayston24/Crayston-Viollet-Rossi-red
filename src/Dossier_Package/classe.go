package Dossier_Package

func (p *Personnage) Tank() {
	p.point_de_vie_maximum = 100
	p.point_de_vie_actuel = 100
	p.Skill = append(p.Skill, "coup de bouclier")

}

func (p *Personnage) Attaquant() {
	p.point_de_vie_maximum = 50
	p.point_de_vie_actuel = 50
	p.Skill = append(p.Skill, "coup de épée")
}

func (p *Personnage) Equilibré() {
	p.point_de_vie_maximum = 75
	p.point_de_vie_actuel = 75
	p.Skill = append(p.Skill, "coup de poing")
}
