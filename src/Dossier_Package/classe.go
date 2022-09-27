package Dossier_Package

func (p *Personnage) Tank() {
	p.classe = "tank"
	p.point_de_vie_maximum = 100
	p.point_de_vie_actuel = 100
	p.mana_actuel = 30
	p.mana_maximum = 30
	p.Skill = append(p.Skill, "coup de bouclier")
	p.addInventory("potion")
	p.niveau = 1
	p.monnaie = 100
	p.taille_inventaire = 10

}

func (p *Personnage) Attaquant() {
	p.classe = "attaquant"
	p.point_de_vie_maximum = 50
	p.point_de_vie_actuel = 50
	p.mana_actuel = 30
	p.mana_maximum = 30
	p.Skill = append(p.Skill, "coup de épée")
	p.addInventory("potion")
	p.niveau = 1
	p.monnaie = 100
	p.taille_inventaire = 10
}

func (p *Personnage) Equilibré() {
	p.classe = "equilibré"
	p.point_de_vie_maximum = 75
	p.point_de_vie_actuel = 75
	p.mana_actuel = 30
	p.mana_maximum = 30
	p.Skill = append(p.Skill, "coup de poing")
	p.addInventory("potion")
	p.niveau = 1
	p.monnaie = 100
	p.taille_inventaire = 10
}
