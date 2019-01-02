package Tennis

type TennisGame struct {
	aScore int
	bScore int
}

var Score map[int]string = getScoreName()

func (t *TennisGame) GetScore() string {
	if t.sameScore() {
		if t.isDeuce() {
			return "Deuce"
		}
		return Score[t.aScore] + " All"
	} else if t.enoughScore() {
		if adv(t.aScore-t.bScore) == 1 {
			return t.getSupperiorName() + " Adv"
		}
		return t.getSupperiorName() + " Win"
	}
	return Score[t.aScore] + " " + Score[t.bScore]
}
func (t *TennisGame) getSupperiorName() string {
	if t.aScore > t.bScore {
		return "Player One"
	} else {
		return "Player Two"
	}
}
func (t *TennisGame) enoughScore() bool {
	if t.aScore >= 4 || t.bScore >= 4 {
		return true
	} else {
		return false
	}
}
func (t *TennisGame) isDeuce() bool {
	if t.aScore >= 3 {
		return true
	} else {
		return false
	}
}
func (t *TennisGame) sameScore() bool {
	if t.aScore == t.bScore {
		return true
	} else {
		return false
	}
}
func (t *TennisGame) setPlayerTwoScore(score int) {
	t.bScore = score
}
func (t *TennisGame) setPlayerOneScore(score int) {
	t.aScore = score
}
func NewTennisGame() *TennisGame {
	return &TennisGame{}
}
func getScoreName() map[int]string {
	return map[int]string{
		0: "Love",
		1: "Fifteen",
		2: "Thirty",
		3: "Forty",
	}
}
func adv(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
