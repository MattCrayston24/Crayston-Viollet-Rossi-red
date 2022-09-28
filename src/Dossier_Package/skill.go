package Dossier_Package

import (
	"fmt"
)

func (p *Personnage) addSkill(str, str1 string, nb, nbr int) {
	if len(p.skill.nom) < 5 {
		p.skill.element = append(p.skill.element, str)
		p.skill.nom = append(p.skill.nom, str1)
		p.skill.degat = append(p.skill.degat, nb)
		p.skill.cout_mana = append(p.skill.cout_mana, nbr)
	}
}

type Skill struct {
	element   []string
	cout_mana []int
	degat     []int
	nom       []string
}

func (p Personnage) Menu_skill_choix() {
		var choix_menu int 
		fmt.Println("______________________________________")
		fmt.Println("|                                    |")
		fmt.Println("|         Menu Skill Choix           |")
		fmt.Println("|____________________________________|")
		fmt.Println("|                                    |")
		fmt.Println("|             1.Feu                  |")
		fmt.Println("|                                    |")
		fmt.Println("|             2.Eau                  |")
		fmt.Println("|                                    |")
		fmt.Println("|             3.Air                  |")
		fmt.Println("|                                    |")
		fmt.Println("|             4.Terre                |")
		fmt.Println("|                                    |")
		fmt.Println("|             5.Quitter              |")
		fmt.Println("|____________________________________|")
		fmt.Scan(&choix_menu)
		for i:= 0;i< 2;i++{
		switch choix_menu {
		case 1:
			p.Menu_skill_Feu()
			break
		case 2:
			p.Menu_skill_Eau()
			break
		case 3:
			p.Menu_skill_Air()
			break
		case 4:
			p.Menu_skill_Terre()
			break
		case 5:
			p.Menu()	
		default:
			fmt.Println("Vous n'avez pas selectionner un reponse valide")
			p.Menu_skill_choix()
		}
	}
}

func (p Personnage) Menu_skill_Feu() {
	var choix_menu int
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|           Menu Skill Feu           |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	fmt.Println("|             1. boule de Feu        |")
	fmt.Println("|                                    |")
	fmt.Println("|             2. esprit de feu       |")
	fmt.Println("|                                    |")
	fmt.Println("|             3. Epée de feu         |")
	fmt.Println("|                                    |")
	fmt.Println("|             4. Epée enflammée      |")
	fmt.Println("|                                    |")
	fmt.Println("|             5. Flèche enflammée    |")
	fmt.Println("|                                    |")
	fmt.Println("|             6. Arc enflammé        |")
	fmt.Println("|                                    |")
	fmt.Println("|             7 .Quitter             |")
	fmt.Println("|____________________________________|")
	fmt.Scan(&choix_menu)

	switch choix_menu {
	case 1:
		if p.verif_classe() == "mage" && p.point_skill > 0 && !p.verif_skill("Boule de Feu") {
			p.addSkill("Feu", "Boule de Feu", 8, 12)
			fmt.Println("Vous avec acquis le skill Boule de Feu")
			p.point_skill --
		} else {
			fmt.Println("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Feu()
	case 2:
		if p.verif_classe() == "mage" && p.point_skill > 0 && !p.verif_skill("esprit de feu"){
			p.addSkill("Feu", "esprit de feu", 10, 15)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill esprit de feu")
		} else {
			fmt.Println("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Feu()
	case 3:
		if p.verif_classe() == "épéiste" && p.point_skill > 0 && !p.verif_skill("Epée de feu"){
			p.addSkill("Feu", "Epée de feu", 8, 9)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill Epée de feu")
		} else {
			fmt.Println("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Feu()
	case 4:
		if p.verif_classe() == "épéiste" && p.point_skill > 0 && !p.verif_skill("Epée enflammée"){
			p.addSkill("Feu", "Epée enflammée", 6, 5)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill Epée enflammée")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Feu()
	case 5:
		if p.verif_classe() == "archer" && p.point_skill > 0 && !p.verif_skill("Flèche enflammée"){
			p.addSkill("Feu", "Flèche enflammée", 7, 6)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill Flèche enflammée")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Feu()
	case 6:
		if p.verif_classe() == "archer" && p.point_skill > 0 && !p.verif_skill("Arc enflammé") {
			p.addSkill("Feu", "Arc enflammé", 10, 14)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill Arc enflammé")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Feu()
	case 7:
		p.Menu_skill_choix()
	default:
		fmt.Println("Vous n'avez pas selectionné une reponse valide")
		p.Menu_skill_Feu()
	}
}

func (p Personnage) Menu_skill_Eau() {
	var choix_menu int
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|           Menu Skill Eau           |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	fmt.Println("|             1. vague géante        |")
	fmt.Println("|                                    |")
	fmt.Println("|             2. mur d'eau           |")
	fmt.Println("|                                    |")
	fmt.Println("|             3. Epée d'Eau          |")
	fmt.Println("|                                    |")
	fmt.Println("|             4. Epée lance d'eau    |")
	fmt.Println("|                                    |")
	fmt.Println("|             5. Flèche d'eau        |")
	fmt.Println("|                                    |")
	fmt.Println("|             6. Arc impermeable     |")
	fmt.Println("|                                    |")
	fmt.Println("|             7 .Quitter             |")
	fmt.Println("|____________________________________|")
	fmt.Scan(&choix_menu)

	switch choix_menu {
	case 1:
		if p.verif_classe() == "mage" && p.point_skill > 0 && !p.verif_skill("vague géante"){
			p.addSkill("Eau", "vague géante", 8, 12)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill vague géante")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Eau()
	case 2:
		if p.verif_classe() == "mage" && p.point_skill > 0 && !p.verif_skill("mur d'eau"){
			p.addSkill("Eau", "mur d'eau", 10, 15)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill mur d'eau")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Eau()
	case 3:
		if p.verif_classe() == "épéiste" && p.point_skill > 0 && !p.verif_skill("Epée d'Eau"){
			p.addSkill("Eau", "Epée d'Eau", 8, 9)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill Epée d'Eau")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Eau()
	case 4:
		if p.verif_classe() == "épéiste" && p.point_skill > 0 && !p.verif_skill("Epée lance d'eau"){
			p.addSkill("Eau", "Epée lance d'eau", 6, 5)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill Epée lance d'eau")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Eau()
	case 5:
		if p.verif_classe() == "archer" && p.point_skill > 0 && !p.verif_skill("Flèche d'eau"){
			p.addSkill("Eau", "Flèche d'eau", 7, 6)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill Flèche d'eau")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Eau()
	case 6:
		if p.verif_classe() == "archer" && p.point_skill > 0 && !p.verif_skill("Arc impermeable"){
			p.addSkill("Eau", "Arc impermeable", 10, 14)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill Arc impermeable")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Eau()
	case 7:
		p.Menu_skill_choix()
	default:
		fmt.Println("Vous n'avez pas selectionné une reponse valide")
		p.Menu_skill_Eau()
	}
}

func (p Personnage) Menu_skill_Air() {
	var choix_menu int
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|           Menu Skill Air           |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	fmt.Println("|             1. tornade             |")
	fmt.Println("|                                    |")
	fmt.Println("|             2. souffle ultime      |")
	fmt.Println("|                                    |")
	fmt.Println("|             3. Epée tornade        |")
	fmt.Println("|                                    |")
	fmt.Println("|             4. lame invisible      |")
	fmt.Println("|                                    |")
	fmt.Println("|             5. Flèche légère       |")
	fmt.Println("|                                    |")
	fmt.Println("|             6. Arc teleportant     |")
	fmt.Println("|                                    |")
	fmt.Println("|             7 .Quitter             |")
	fmt.Println("|____________________________________|")
	fmt.Scan(&choix_menu)

	switch choix_menu {
	case 1:
		if p.verif_classe() == "mage" && p.point_skill > 0 && !p.verif_skill("tornade"){
			p.addSkill("Air", "tornade", 8, 12)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill tornade")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Air()
	case 2:
		if p.verif_classe() == "mage" && p.point_skill > 0 && !p.verif_skill("souffle ultime") {
			p.addSkill("Air", "souffle ultime", 10, 15)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill souffle ultime")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Air()
	case 3:
		if p.verif_classe() == "épéiste" && p.point_skill > 0 && !p.verif_skill("Epée tornade"){
			p.addSkill("Air", "Epée tornade", 8, 9)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill Epée tornade")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Air()
	case 4:
		if p.verif_classe() == "épéiste" && p.point_skill > 0 && !p.verif_skill("lame invisible"){
			p.addSkill("Air", "lame invisible", 6, 5)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill lame invisible")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Air()
	case 5:
		if p.verif_classe() == "archer" && p.point_skill > 0 && !p.verif_skill("Flèche légère"){
			p.addSkill("Air", "Flèche légère", 7, 6)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill Flèche légère")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Air()

	case 6:
		if p.verif_classe() == "archer" && p.point_skill > 0 && !p.verif_skill("Arc teleportant"){
			p.addSkill("Air", "Arc teleportant", 10, 14)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill Arc teleportant")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Air()
	case 7:
		p.Menu_skill_choix()
	default:
		fmt.Println("Vous n'avez pas selectionné une reponse valide")
		p.Menu_skill_Air()
	}
}

func (p Personnage) Menu_skill_Terre() {
	var choix_menu int
	fmt.Println("______________________________________")
	fmt.Println("|                                    |")
	fmt.Println("|          Menu Skill Terre          |")
	fmt.Println("|____________________________________|")
	fmt.Println("|                                    |")
	fmt.Println("|             1. racine tueuse       |")
	fmt.Println("|                                    |")
	fmt.Println("|             2. sable mouvant       |")
	fmt.Println("|                                    |")
	fmt.Println("|             3. Epée en pierre      |")
	fmt.Println("|                                    |")
	fmt.Println("|             4. lame en diamant     |")
	fmt.Println("|                                    |")
	fmt.Println("|             5. Flèche en or        |")
	fmt.Println("|                                    |")
	fmt.Println("|             6. Arc en marbre       |")
	fmt.Println("|                                    |")
	fmt.Println("|             7 .Quitter             |")
	fmt.Println("|____________________________________|")
	fmt.Scan(&choix_menu)

	switch choix_menu {
	case 1:
		if p.verif_classe() == "mage" && p.point_skill > 0 && !p.verif_skill("racine tueuse"){
			p.addSkill("Terre", "racine tueuse", 8, 12)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill racine tueuse")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Terre()
	case 2:
		if p.verif_classe() == "mage" && p.point_skill > 0 && !p.verif_skill("sable mouvant"){
			p.addSkill("Terre", "sable mouvant", 10, 15)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill sable mouvant")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Terre()
	case 3:
		if p.verif_classe() == "épéiste" && p.point_skill > 0 && !p.verif_skill("Epée en pierre"){
			p.addSkill("Terre", "Epée en pierre", 8, 9)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill Epée en pierre")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Terre()
	case 4:
		if p.verif_classe() == "épéiste" && p.point_skill > 0&& !p.verif_skill("lame en diamant") {
			p.addSkill("Terre", "lame en diamant", 6, 5)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill lame en diamant")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Terre()
	case 5:
		if p.verif_classe() == "archer" && p.point_skill > 0&& !p.verif_skill("Flèche en or") {
			p.addSkill("Terre", "Flèche en or", 7, 6)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill Flèche en or")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Terre()
	case 6:
		if p.verif_classe() == "archer" && p.point_skill> 0 && !p.verif_skill("Arc en marbre") {
			p.addSkill("Terre", "Arc en marbre", 10, 14)
			p.point_skill --
			fmt.Println("Vous avec acquis le skill Arc en marbre")
		} else {
			fmt.Print("Vous n'avez pas accés à ce skill ou vous n'avez pas assez de point de skill")
			fmt.Println("Vous avez ", p.point_skill, "point de skill")
		}
		p.Menu_skill_Terre()
	case 7:
		p.Menu_skill_choix()
	default:
		fmt.Println("Vous n'avez pas selectionné une reponse valide")
		p.Menu_skill_Terre()
	}
}