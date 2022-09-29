package Dossier_Package

import "fmt"

func (p *Personnage) verif_materiaux(nb int, str string) bool {
	count := 0
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == str {
			count++
		}
	}
	return count >= nb
}

func (p Personnage) CheckInventory() bool {

	return len(p.inventaire) <= p.taille_inventaire
}

func (p *Personnage) addInventory(s string) {
	if len(p.inventaire) < p.taille_inventaire {
		p.inventaire = append(p.inventaire, s)
	}
}
func (p *Personnage) removeInventory(s string) {
	for i :=0;i<len(p.inventaire);i++{
		if p.inventaire[i] == s {
			p.inventaire = append(p.inventaire[:i], p.inventaire[i+1:]...)
		}
	}
}

func (p Personnage) AccèsInventaire() {
	p.Menu_Inventaire()
	fmt.Println(p.inventaire)
}
