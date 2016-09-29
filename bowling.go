package bowling

import "fmt"


type Frame struct {
	firstThrow  int
	secondThrow int
}

func GetScore(game []Frame) (int, error) {
	score := 0

	if(len(game)!= 10){
		return 0,fmt.Printf("number of tuplet should be ten")

	}else{
		for i := 0; i < len(game); i++ {
			if(game[i].firstThrow<0 || game[i].secondThrow <0){
				return 0,nil
			}else{
					totalUpplet :=  game[i].firstThrow + game[i].secondThrow
				if( totalUpplet > 10){
					return 0,nil
				}else{

					if(totalUpplet== 10 && i+1 < len(game)){

						if game[i].firstThrow == 10{

							score = score + totalUpplet +game[i+1].firstThrow + game[i+1].secondThrow
						}else{
							score = score + totalUpplet +game[i+1].firstThrow 
						}

						
						
					}else{
						score = score + totalUpplet 
					}

					

				}

			}
		
		}
	return score, nil
	}
	
}
