package Dossier_Package

func (p Personnage) verif_materiaux(nb int, str string) bool {
	count := 0
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == str {
			count++
		}
	}
	return count == nb
}

func (p Personnage) CheckInventory() bool {

	return len(p.inventaire) != 10
}

func (p *Personnage) addInventory(s string) {
	if len(p.inventaire) < 10 {
		p.inventaire = append(p.inventaire, s)
	}
}
func (p *Personnage) removeInventory(s string) {
	for i := range s {
		if p.inventaire[i] == s {
			p.inventaire = append(p.inventaire[:i], p.inventaire[i+1:]...)
		}
	}
}
