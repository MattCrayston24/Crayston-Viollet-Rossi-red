func (p *Personnage) PoisonPot(){
	p.point_de_vie_actuel -= 10
	time.Sleep(1 * time.Second)
	p.point_de_vie_actuel -= 10
	time.Sleep(1 * time.Second)
	p.point_de_vie_actuel -= 10
}