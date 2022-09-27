package Dossier_Package

func (p *Personnage) Archer() {
	p.classe = "Archer"
	p.point_de_vie_maximum = 80
	p.point_de_vie_actuel = 80
	p.Skill = append(p.Skill, "Flèche charger")
	p.taille_inventaire = 10
	p.addInventory("potion")
	p.niveau = 1
	p.monnaie = 100
	p.initiative = 12
	p.mana_actuel = 20
	p.mana_maximum = 20
	p.points_attaque = 5
	p.mon_equipement.armes = "Arc en d'entrainement"

}

func (p *Personnage) Epéiste() {
	p.classe = "épéiste"
	p.point_de_vie_maximum = 100
	p.point_de_vie_actuel = 100
	p.Skill = append(p.Skill, "Coup vertical")
	p.taille_inventaire = 10
	p.addInventory("potion")
	p.niveau = 1
	p.monnaie = 100
	p.initiative = 8
	p.mana_actuel = 15
	p.mana_maximum = 15
	p.points_attaque = 5
	p.mon_equipement.armes = "épée d'entrainement"
}

func (p *Personnage) Mage() {
	p.classe = "mage"
	p.point_de_vie_maximum = 70
	p.point_de_vie_actuel = 70
	p.Skill = append(p.Skill, "Boule de feu")
	p.taille_inventaire = 10
	p.addInventory("potion")
	p.niveau = 1
	p.monnaie = 100
	p.initiative = 15
	p.mana_actuel = 50
	p.mana_maximum = 50
	p.points_attaque = 5
	p.mon_equipement.armes = "Baton d'entrainement "
}

func (p Personnage) verif_classe() string {
	if p.classe == "archer" {
		return "archer"
	} else if p.classe == "épéiste" {
		return "épéiste"
	} else if p.classe == "mage" {
		return "mage"
	}
	return "aucune classe"
}
