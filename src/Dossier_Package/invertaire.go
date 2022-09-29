package Dossier_Package

import "fmt"

//Fonction qui verifie si les matériaux sont dans l'inventaire
func (p *Personnage) verif_materiaux(nb int, str string) bool {
	count := 0
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == str {
			count++
		}
	}
	return count >= nb
}

//Fonction qui permet de vérifier si y as assez de place dans l'inventaire
func (p Personnage) CheckInventory() bool {

	return len(p.inventaire) <= p.taille_inventaire
}

//Fonction qui permet d'ajouter un objet dans l'inventaire
func (p *Personnage) addInventory(s string) {
	if len(p.inventaire) < p.taille_inventaire {
		p.inventaire = append(p.inventaire, s)
	}
}

//Fonction qui permet de supprimer un objet dans l'inventaire
func (p *Personnage) removeInventory(s string) {
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == s {
			p.inventaire = append(p.inventaire[:i], p.inventaire[i+1:]...)
		}
	}
}

//Fonction qui permet de print l'inventaire
func (p Personnage) AccèsInventaire() {

	p.Menu_Inventaire()
	if len(p.inventaire) > 0 {
		fmt.Println(p.inventaire)
	} else {
		fmt.Println("votre inventaire est vide")
	}
}
