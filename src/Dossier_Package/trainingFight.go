func (p *Personnage, m *monstre)trainingFight(str string){
	nbtours:= []int
	var attaque_choisi string
	attaquemonstre := 0
	for i:=0; pv_joueur <= 0 || pv_monstre <= 0; i++{
		if i%2 == 0{
			fmt.Println("Pour attaquer avec une lance taper 1, \n Pour attaquer avec l'épée taper 2")
			fmt.Scan(attaque_choisi)
			if attaque_choisi == "1"{
				m.point_de_vie_actuel = m.point_de_vie_actuel - 30
			}
			if attaque_choisi == "2"{
				m.point_de_vie_actuel = m.point_de_vie_actuel - 25
				if m.point_de_vie_shield > 10{
					m.point_de_vie_shield = m.point_de_vie_shield - 10
				}else{
					m.point_de_vie_shield = m.point_de_vie_shield 
				}
			}
			nbtours += 1
		}else{
			if attaquemonstre <= 2{
				p.point_de_vie_actuel = p.point_de_vie_actuel - 25
				attaquemontre +=1
				fmt.Println("Le niveau du montre est actuellement 1")
				nbtours += 1
			}
			if attaquemontre <= 4{
				p.point_de_vie_actuel = p.point_de_vie_actuel - 20
				if p.point_de_vie_shield > 10{
					p.point_de_vie_shield = p.point_de_vie_shield - 10
				}else{
					p.point_de_vie_actuel = p.point_de_vie_actuel - 10
				}
				fmt.Println("Le niveau du montre est actuellement 2")
				nbtours += 1
			}
			
			if attaquemonstre <= 6{
				p.point_de_vie_actuel = p.point_de_vie_actuel - 35
				fmt.Println("Le niveau du montre est actuellement 3")
				nbtours += 1
			}
			
			if attaquemonstre <= 8{
				p.point_de_vie_actuel = p.point_de_vie_actuel - 40
				if p.point_de_vie_shield > 15{
					p.point_de_vie_shield = p.point_de_vie_shield - 15
				}else{
					p.point_de_vie_actuel = p.point_de_vie_actuel - 10
				}
				fmt.Println("Le niveau du montre est actuellement 4")
				nbtours += 1
			}

		}
		print(nbtours)
	}
}