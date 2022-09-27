package Dossier_Package

type Skill struct{
	feu []string
	eau []string
	air []string
}

func (p *Personnage) addSkill(str string) {
	feu := ["boule de feu","poing de feu","esprit de feu"]
	eau := ["vague titanesque","tsunami"]
	air := ["tornade","tempête"]
	var skill_choisi int
	fmt.Print("Si vous souhaitez choisir le skill feu taper 1 \n le skill eau 2 \n le skill air 3")
	fmt.Scan(&skill_choisi)
	switch skill_choisi{
	case 1:
		var type_choisi int
		fmt.Print("boule de feu : 1 \n poing de feu : 2 \n esprit de feu : 3")
		fmt.Scan(&type_choisi)
		switch type_choisi{
		case 1:
			if len(p.Skill) < 5 {
				p.Skill = append(p.Skill, "boule de feu")
			}else{
				fmt.Print("Vous avez déjà trop de skills")
			}
		case 2:
			if len(p.Skill) < 5 {
				p.Skill = append(p.Skill, "poing de feu")
			}else{
				fmt.Print("Vous avez déjà trop de skills")
			}
		case 3:
			if len(p.Skill) < 5 {
				p.Skill = append(p.Skill, "esprit de feu")
			}else{
				fmt.Print("Vous avez déjà trop de skills")
			}
		}
	case 2:
		var type_choisi int
		fmt.Print("vague titanesque : 1 \n tsunami : 2")
		fmt.Scan(&type_choisi)
		switch type_choisi{
		case 1:
			if len(p.Skill) < 5 {
				p.Skill = append(p.Skill, "vague titanesque")
			}else{
				fmt.Print("Vous avez déjà trop de skills")
			}
		case 2:
			if len(p.Skill) < 5 {
				p.Skill = append(p.Skill, "tsunami")
			}else{
				fmt.Print("Vous avez déjà trop de skills")
			}
		}
	case 3:
		var type_choisi int
		fmt.Print(" tornade : 1 \n tempête : 2")
		fmt.Scan(&type_choisi)
		switch type_choisi{
		case 1:
			if len(p.Skill) < 5 {
				p.Skill = append(p.Skill, "tornade")
			}else{
				fmt.Print("Vous avez déjà trop de skills")
			}
		case 2:
			if len(p.Skill) < 5 {
				p.Skill = append(p.Skill, "tempête")
			}else{
				fmt.Print("Vous avez déjà trop de skills")
			}

	}
}

