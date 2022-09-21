package Dossier_Package

func (p *Personnage) Tank() {
	p.classe = "tank"
	p.point_de_vie_maximum = 100
	p.point_de_vie_actuel = 100
	p.Skill = append(p.Skill, "coup de bouclier")
	p.addInventory("potion")
	p.niveau = 1
	p.monnaie = 100

}

func (p *Personnage) Attaquant() {
	p.classe = "attaquant"
	p.point_de_vie_maximum = 50
	p.point_de_vie_actuel = 50
	p.Skill = append(p.Skill, "coup de épée")
	p.addInventory("potion")
	p.niveau = 1
	p.monnaie = 100
}

func (p *Personnage) Equilibré() {
	p.classe = "equilibré"
	p.point_de_vie_maximum = 75
	p.point_de_vie_actuel = 75
	p.Skill = append(p.Skill, "coup de poing")
	p.addInventory("potion")
	p.niveau = 1
	p.monnaie = 100
}
