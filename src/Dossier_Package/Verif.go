package Dossier_Package

import (
	"unicode"
	"bufio"
	"os"
)	
func (p *Personnage) verif_nom(nom string) bool {
	var nb_lettre int
	for _, letter := range nom {
		if unicode.IsLetter(letter) {
			nb_lettre += 1
		}
	}
	return nb_lettre == len(nom)
}

func (p *Personnage) majuscule(nom string) string {
	var nom_maj string
	for i, letter := range nom {
		if i == 0 {
			nom_maj += string(unicode.ToUpper(letter))
		} else {
			nom_maj += string(unicode.ToLower(letter))
		}
	}
	return nom_maj
}

func verif_element_faible(str, str1 string) bool {
	if str != str1 {
		if str == "feu" && str1 == "eau" || str == "eau" && str1 == "terre" || str == "terre" && str1 == "air" || str == "air" && str1 == "feu" {
			return true
		} else {
			return false
		}
	}
	return false
}

func verif_element_fort(str, str1 string) bool {
	if str != str1 {
		if str == "eau" && str1 == "feu" || str == "terre" && str1 == "eau" || str == "air" && str1 == "terre" || str == "feu" && str1 == "air" {
			return true
		}
	} else {
		return false
	}
	return false
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

func (p Personnage) verif_monnaie(nb int) bool {
	return p.monnaie >= nb
}

func (p *Personnage) verif_skill(str string) bool {
	for i := 0; i < len(p.skill.nom); i++ {
		if p.skill.nom[i] == str {
			return true
		}
	}
	return false
}

func (p *Personnage)Scan() string{
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
