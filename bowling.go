package bowling



type Frame struct {
	firstThrow  int
	secondThrow int
}

func GetScore(game []Frame) (int, error) {
	score := 0

	if(len(game)!= 10){
		return 0,fmt.Errorf("la taille du Tuple est différente de 10 ") 

	}else{
		for i := 0; i < len(game); i++ {
			if(game[i].firstThrow<0 || game[i].secondThrow <0){
				
				return 0,fmt.Errorf("valeur du tuple négative non souhaitée")
			}else{
					totalUpplet :=  game[i].firstThrow + game[i].secondThrow
				
				if( totalUpplet > 10){
					
					return 0,fmt.Errorf("Tuple supérieur à 10")
				}else{

					if(totalUpplet== 10 && i+1 < len(game)){

						if game[i].firstThrow == 10{
							//  le cas du Strike 
							score = score + totalUpplet +game[i+1].firstThrow + game[i+1].secondThrow
						}else{
							//  le cas du Spare
							score = score + totalUpplet +game[i+1].firstThrow 
						}

						
						
					}else{
						score = score + totalUpplet 
					}

					

				}

			}
		
		}
	return score,fmt.Errorf("valeurs du tuple négatives")
	}
	
}
