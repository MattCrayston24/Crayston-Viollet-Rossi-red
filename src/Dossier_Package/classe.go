package Dossier_Package

func (p *Personnage) Archer() {
	p.classe = "archer"
	p.point_de_vie_maximum = 80
	p.point_de_vie_actuel = 80
	p.taille_inventaire = 50
	p.inventaire = []string{"Potion de Vie", "arc d'entrainement", "plastron en cuir"}
	p.niveau = 1
	p.monnaie = 100
	p.initiative = 12
	p.mana_actuel = 20
	p.mana_maximum = 20
	p.attaque_base = 5
	p.point_skill = 1
}

func (p *Personnage) Epéiste() {
	p.classe = "épéiste"
	p.point_de_vie_maximum = 100
	p.point_de_vie_actuel = 100
	p.taille_inventaire = 50
	p.inventaire = []string{"Potion de Vie", "épée d'entrainement"}
	p.niveau = 1
	p.monnaie = 100
	p.initiative = 8
	p.mana_actuel = 15
	p.mana_maximum = 15
	p.attaque_base = 5
	p.point_skill = 1
}

func (p *Personnage) Mage() {
	p.classe = "mage"
	p.point_de_vie_maximum = 70
	p.point_de_vie_actuel = 70
	p.taille_inventaire = 50
	p.inventaire = []string{"Potion de Vie", "Baton d'entrainement "}
	p.niveau = 1
	p.monnaie = 100
	p.initiative = 15
	p.mana_actuel = 50
	p.mana_maximum = 50
	p.attaque_base = 5
	p.point_skill = 1
}
