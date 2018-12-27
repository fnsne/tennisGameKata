package tennis

type TennisGame struct {
	aScore    int
	bScore    int
	scoreName map[int]string
	aName     string
	bName     string
}

func (t *TennisGame) setPlayerAScore(score int) {
	t.aScore = score
}
func (t *TennisGame) setPlayerBScore(score int) {
	t.bScore = score
}
func (t *TennisGame) GetScore() string {
	if t.sameScore() {
		if t.aScore < 3 {
			return t.scoreName[t.aScore] + " All"
		} else {
			return "Deuce"
		}
	} else {
		if t.enoughScore() {
			if abs(t.aScore-t.bScore) == 1 {
				return t.superiorPlayer() + " Adv"
			} else {
				return t.superiorPlayer() + " Win"
			}
		}
		return t.scoreName[t.aScore] + " " + t.scoreName[t.bScore]
	}
}
func (t *TennisGame) enoughScore() bool {
	return t.aScore >= 3 && t.bScore >= 3
}
func (t *TennisGame) superiorPlayer() string {
	if t.aScore > t.bScore {
		return t.aName
	} else {
		return t.bName
	}
}
func (t *TennisGame) sameScore() bool {
	if t.aScore == t.bScore {
		return true
	} else {
		return false
	}
}
func getTennisGame() *TennisGame {
	return &TennisGame{
		aScore:    0,
		bScore:    0,
		scoreName: getScoreName(),
		aName:     "Player A",
		bName:     "Player B",
	}
}
func getScoreName() map[int]string {
	scoreName := make(map[int]string)
	scoreName[0] = "Love"
	scoreName[1] = "Fifteen"
	scoreName[2] = "Thirty"
	scoreName[3] = "Forty"
	return scoreName
}
func abs(num int) int {
	if num < 0 {
		return num * -1
	} else {
		return num
	}
}
