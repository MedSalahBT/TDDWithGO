package bowling

type Frame struct {
	firstThrow  int
	secondThrow int
}

func GetScore(game []Frame) (int, error) {
	score := 0
	for i := 0; i < len(game); i++ {
		score = score + score[i].firstThrow + score[i].secondThrow
	}
	return score, nil
}
