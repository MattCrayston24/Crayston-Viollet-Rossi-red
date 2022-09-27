func histoire(){
	var choix_intrigue int
	fmt.Print("Bienvenue dans Ora, un monde dans lequel des monstres dangereux vivent et dans lequel vous devrez vous en sortir, Bonne chance !")
	fmt.Print("Si vous souhaitez connaître l'histoire de votre personnage tapez 1 \n Si vous soiuhaitez partir à l'aventure directement tapez 2 \n Quitter le jeu 3")
	fmt.Scan(&choix_intrigue)
	switch choix_intrigue {
	case 1:
		fmt.Print("Il était une fois un jeune chevalier combattant des dragons a ses heures perdues, fou amoureux de sa femme Dana, 
		un jour un monstre inconnu du monde Ora apparu et devora Dana, ce monstre géant et d'une force spectaculaire fut la raison 
		de vivre de ce chevalier, il était a sa traque pendant si longtemps qu'un jour il se décida a l'éliminer, il engloutit tellement de 
		potions que ses yeux devint noir, il brandit son épée et trancha la tête du monstre pendant un combat sanglant, désormais vous êtes ce 
		chevalier et votre but sera de tuer tout les monstres restant dans ce monde !")
	case 2:
		Menu()
	case 3:
		break
	}
}