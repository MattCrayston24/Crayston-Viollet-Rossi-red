package Dossier_Package

func (p *Personnage) Marchand(s string){ 
	var item_choisi string
	if len(p.inventaire) < 10{
		fmt.Println("Potion de bouclier taper 1 \n Potion de vie taper 2 ")	
		fmt.Scan(&item_choisi)
		if item_choisi == "1"{
			p.inventaire = append(p.inventaire, "Potion de bouclier")
			p.monnaie -= 30
		} else if item_choisi == "2"{
			p.inventaire = append(p.inventaire, "Potion de vie")
			p.monnaie -= 20
			}
	}
}