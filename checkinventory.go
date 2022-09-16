func (p *Personnage)CheckInventory()bool{
	if len(p.inventaire) != 10{
		return true
	}
	return false
}