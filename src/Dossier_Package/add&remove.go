func (p *Personnage) addInventory(s string){
	if len(p.inventaire) < 10{
		p.inventaire = append(p.inventaire, s)
	}
}
func (p *Personnage) removeInventory(s string){
	for i,_:= range s{
		if p.inventaire[i] == s{
			p.inventaire = append(p.inventaire[:i], p.inventaire[i+1:]...)
		} 
	}
}