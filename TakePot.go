import "fmt"

func (p *Personnage) TakePot(){
	potion := 20
	for i,_ := range p.inventaire {
		if p.inventaire[i] == "potion"{
			if p.point_de_vie_actuel + potion > p.point_de_vie_maximum{
				p.point_de_vie_actuel = p.point_de_vie_maximum
				p.inventaire = append(p.inventaire[:i],p.inventaire[i+1:]...)
				break
			}else{
				p.point_de_vie_actuel += potion
				p.inventaire = append(p.inventaire[:i],p.inventaire[i+1:]...)
				break
			}
		}	
	}	
}