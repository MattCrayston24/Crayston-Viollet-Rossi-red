package Dossier_Package

type Equipement struct {
	casque              string
	plastron            string
	bottes              string
	armes               string
	liste_casque        []string
	liste_plastron      []string
	liste_bottes        []string
	liste_armes_archer  []string
	liste_armes_épéiste []string
	liste_armes_mage    []string
}

func (e *Equipement) Init_List() {
	e.liste_casque = []string{"casque en cuir", "casque en fer", "casque en mithril"}
	e.liste_plastron = []string{"plastron en cuir", "plastron en fer", "plastron en mithril"}
	e.liste_bottes = []string{"bottes en cuir", "bottes en fer", "bottes en mithril"}
	e.liste_armes_archer = []string{"arc en bois", "arc en bois travailler", "arc en acier", "arc en mitrhil"}
	e.liste_armes_épéiste = []string{"épée d'entrainement", "épée en fer", "epée en acier", "épée en mitrhil"}
	e.liste_armes_mage = []string{"baton d'entrainement", "baton en fer  ", "baton en acier ", "épée en mithril"}
}

func (p *Personnage) Ajout_Stat_equipement_casque() {
	casque := p.mon_equipement.casque

	switch casque {
	case p.mon_equipement.liste_casque[0]:
		p.point_de_vie_maximum += 10
	case p.mon_equipement.liste_casque[1]:
		p.point_de_vie_maximum += 17
	case p.mon_equipement.liste_casque[2]:
		p.point_de_vie_maximum += 24
	default:
		break
	}

}

func (p *Personnage) Ajout_Stat_equipement_plastron() {
	plastron := p.mon_equipement.plastron

	switch plastron {
	case p.mon_equipement.liste_plastron[0]:
		p.point_de_vie_maximum += 20
	case p.mon_equipement.liste_plastron[1]:
		p.point_de_vie_maximum += 27
	case p.mon_equipement.liste_plastron[2]:
		p.point_de_vie_maximum += 34
	default:
		break
	}

}
func (p *Personnage) Ajout_Stat_equipement_bottes() {
	bottes := p.mon_equipement.bottes

	switch bottes {
	case p.mon_equipement.liste_bottes[0]:
		p.point_de_vie_maximum += 7
	case p.mon_equipement.liste_bottes[1]:
		p.point_de_vie_maximum += 14
	case p.mon_equipement.liste_bottes[2]:
		p.point_de_vie_maximum += 21
	default:
		break
	}

}
func (p *Personnage) Mettre_casque(nb int, e Equipement) {
	if p.mon_equipement.casque != "" {
		p.addInventory(p.mon_equipement.casque)
		p.mon_equipement.casque = ""
	}
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == e.liste_casque[nb-1] {
			p.mon_equipement.casque = p.inventaire[i]
			p.removeInventory(e.liste_casque[nb-1])
			p.Ajout_Stat_equipement_casque()
		}
	}

}

func (p *Personnage) Mettre_Plastron(nb int, e Equipement) {
	if p.mon_equipement.plastron != "" {
		p.addInventory(p.mon_equipement.plastron)
		p.mon_equipement.plastron = ""
	}
	if nb == 1 {
		for i := 0; i < len(p.inventaire); i++ {
			if p.inventaire[i] == e.liste_plastron[nb-1] {
				p.mon_equipement.plastron = p.inventaire[i]
				p.removeInventory(e.liste_plastron[nb-1])
				p.Ajout_Stat_equipement_plastron()
			}
		}
	}
}

func (p *Personnage) Mettre_bottes(nb int, e Equipement) {
	if p.mon_equipement.bottes != "" {
		p.addInventory(p.mon_equipement.bottes)
		p.mon_equipement.bottes = ""
	}
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == e.liste_bottes[nb-1] {
			p.mon_equipement.bottes = p.inventaire[i]
			p.removeInventory(e.liste_bottes[nb])

		}
	}
}

func (p *Personnage) Mettre_armes_archer(nb int, e Equipement) {
	if p.mon_equipement.armes != "" {
		p.addInventory(p.mon_equipement.armes)
		p.mon_equipement.armes = ""
	}
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == e.liste_armes_archer[nb-1] {
			p.mon_equipement.bottes = p.inventaire[i]
			p.removeInventory(e.liste_armes_archer[nb-1])
			p.Ajout_Stat_armes_archer()

		}
	}
}

func (p *Personnage) Mettre_armes_épéiste(nb int, e Equipement) {
	if p.mon_equipement.armes != "" {
		p.addInventory(p.mon_equipement.armes)
		p.mon_equipement.armes = ""
	}
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == e.liste_armes_épéiste[nb-1] {
			p.mon_equipement.bottes = p.inventaire[i]
			p.removeInventory(e.liste_armes_épéiste[nb-1])
			p.Ajout_Stat_armes_épéiste()

		}
	}
}

func (p *Personnage) Mettre_armes_mage(nb int, e Equipement) {
	if p.mon_equipement.armes != "" {
		p.addInventory(p.mon_equipement.armes)
		p.mon_equipement.armes = ""
	}
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == e.liste_armes_mage[nb-1] {
			p.mon_equipement.bottes = p.inventaire[i]
			p.removeInventory(e.liste_armes_mage[nb-1])
			p.Ajout_Stat_armes_mage()

		}
	}
}

func (p *Personnage) Ajout_Stat_armes_archer() {
	plastron := p.mon_equipement.plastron

	switch plastron {
	case p.mon_equipement.liste_armes_archer[0]:
		p.points_attaque += 4
	case p.mon_equipement.liste_armes_archer[1]:
		p.points_attaque += 7
	case p.mon_equipement.liste_armes_archer[2]:
		p.points_attaque += 10
	default:
		break
	}
}

func (p *Personnage) Ajout_Stat_armes_épéiste() {
	plastron := p.mon_equipement.plastron

	switch plastron {
	case p.mon_equipement.liste_armes_épéiste[0]:
		p.points_attaque += 4
	case p.mon_equipement.liste_armes_épéiste[1]:
		p.points_attaque += 7
	case p.mon_equipement.liste_armes_épéiste[2]:
		p.points_attaque += 10
	default:
		break
	}
}

func (p *Personnage) Ajout_Stat_armes_mage() {
	plastron := p.mon_equipement.plastron

	switch plastron {
	case p.mon_equipement.liste_armes_mage[0]:
		p.points_attaque += 4
	case p.mon_equipement.liste_armes_mage[1]:
		p.points_attaque += 7
	case p.mon_equipement.liste_armes_mage[2]:
		p.points_attaque += 10
	default:
		break
	}
}
