func (p *Personnage) Init(){ 
	var classe_choisi string
	var nom_choisi string
	for{
		fmt.Print("Choisissez votre nom comprenant au moins 3 caractères :")
		fmt.Scan(&nom_choisi)
		if len(nom_choisi) >= 3 {
			fmt.Print("Votre nom est : ", nom_choisi)
			p.nom = nom_choisi
			break
		}
	}
	for{
		fmt.Print("Choisissez votre classe, \n Si vous voulez la classe tank taper 1, \n Si vous voulez la classe attaquant taper 2, \n Si vous voulez la calsse équilibré tapez 3 :")
		fmt.Scan(&classe_choisi)
		if classe_choisi == "1"{
			p.classe = "tank"
			fmt.Print("Votre classe est : ",p.classe)
			break
		}else if classe_choisi == "2"{
			p.classe = "attaquant"
			fmt.Print("Votre classe est : ",p.classe)
			break
		}else if classe_choisi == "3"{
			p.classe = "equilibré"
			fmt.Print("Votre classe est : ",p.classe)
			break
		}
	}
		if p.classe == "tank"{
		p.point_de_vie_maximum = 100
		p.point_de_vie_actuel = 100
	}else if p.classe == "attaquant"{ 
		p.point_de_vie_maximum = 50
		p.point_de_vie_actuel = 50
	}else if p.classe == "equilibré"{ 
		p.point_de_vie_maximum = 75
		p.point_de_vie_actuel = 75
	}
	p.niveau = 1
	p.money = 100
	fmt.Println("\n Vos point de vie maximum sont : ",p.point_de_vie_maximum)
	fmt.Println(" Vos points de vie actuel sont : ",p.point_de_vie_actuel)
	fmt.Println(" Votre niveau actuel est : ",p.niveau)
	fmt.Println(" Vous avez  : ",p.money," rubis")
}