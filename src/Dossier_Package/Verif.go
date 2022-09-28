package Dossier_Package
import "unicode"
func (p *Personnage) verif_nom(nom string) bool {
	var nb_lettre int
	for _, letter := range nom {
		if unicode.IsLetter(letter) == true {
			nb_lettre += 1
		}
	}
	if nb_lettre >= 3 {
		return true
	}
	return false
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
