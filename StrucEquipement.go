type Equipement struct{
	casque string
	plastron string
	bottes string
}

type Personnage struct{
	nom string
	classe string
	niveau int
	point_de_vie_maximum int 
	point_de_vie_actuel int
	inventaire []string
	money int
	mon_equipement	Equipement
}