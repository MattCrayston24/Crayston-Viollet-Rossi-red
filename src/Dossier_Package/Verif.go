package Dossier_Package

import "unicode"

func (p *Personnage) verif_nom(nom string) bool {
	var nb_lettre int
	for _, letter := range nom {
		if unicode.IsLetter(letter) {
			nb_lettre += 1
		}
	}
	return nb_lettre >= 3
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
