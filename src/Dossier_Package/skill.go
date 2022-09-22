func (p *Personnage) addSkill(str string){
	if len(p.Skill) < 5{
		p.Skill = append(p.Skill, str)
	}
}