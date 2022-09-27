package Dossier_Package

type Equipement struct {
	casque   string
	plastron string
	bottes   string
}

func (p *Personnage) Ajout_Stat_equipement() {
	casque := p.mon_equipement.casque
	plastron := p.mon_equipement.plastron
	bottes := p.mon_equipement.bottes

	switch {
	case bottes == "jambière en cuir":
		p.point_de_vie_maximum += 15
	case plastron == "plastron en cuir":
		p.point_de_vie_maximum += 20
	case casque == "casque en cuir":
		p.point_de_vie_maximum += 10
	default:
		break
	}

}

func (p *Personnage) Mettre_casque(nb int) {
	if p.mon_equipement.casque != "" {
		p.addInventory(p.mon_equipement.casque)
		p.mon_equipement.casque = ""
	}
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == "casque de cuir" {
			p.mon_equipement.casque = p.inventaire[i]
			p.removeInventory("casque de cuir")
		}
	}

}

func (p *Personnage) Mettre_Plastron(nb int) {
	if p.mon_equipement.plastron != "" {
		p.addInventory(p.mon_equipement.plastron)
		p.mon_equipement.plastron = ""
	}
	if nb == 1 {
		for i := 0; i < len(p.inventaire); i++ {
			if p.inventaire[i] == "plastron de cuir " {
				p.mon_equipement.plastron = p.inventaire[i]
				p.removeInventory("plastron de cuir ")
			}
		}

	}
}

func (p *Personnage) Mettre_bottes(nb int) {
	if p.mon_equipement.bottes != "" {
		p.addInventory(p.mon_equipement.bottes)
		p.mon_equipement.bottes = ""
	}
	if nb == 1 {
		for i := 0; i < len(p.inventaire); i++ {
			if p.inventaire[i] == "jambière de cuir" {
				p.mon_equipement.bottes = p.inventaire[i]
				p.removeInventory("jambière de cuir")

			}
		}
	}
}
