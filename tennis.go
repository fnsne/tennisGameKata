package tennis

type Tennis struct {
	aScore int
	bScore int
}

var Score map[int]string = getScoreName()

func (t *Tennis) GetScore() string {
	if t.sameScore() {
		if t.isDeuce() {
			return "Deuce"
		} else {
			return Score[t.aScore] + " All"
		}
	} else {
		if t.enoughScore() {
			if Adv(t.aScore-t.bScore) == 1 {
				return t.superiorUser() + " Adv"
			} else {
				return t.superiorUser() + " Win"
			}
		}
		return Score[t.aScore] + " " + Score[t.bScore]
	}
}
func (t *Tennis) isDeuce() bool {
	if t.aScore >= 3 {
		return true
	} else {
		return false
	}
}
func (t *Tennis) sameScore() bool {
	if t.aScore == t.bScore {
		return true
	} else {
		return false
	}
}
func (t *Tennis) superiorUser() string {
	if t.aScore > t.bScore {
		return "Player One"
	} else {
		return "Player Two"
	}
}
func Adv(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
func (t *Tennis) enoughScore() bool {
	if t.aScore >= 3 && t.bScore >= 3 {
		return true
	} else {
		return false
	}
}
func getScoreName() map[int]string {
	return map[int]string{
		0: "Love",
		1: "Fifteen",
		2: "Thirty",
		3: "Forty",
	}
}
func (t *Tennis) setPlayerTwoScore(score int) {
	t.bScore = score
}
func (t *Tennis) setPlayerOneScore(score int) {
	t.aScore = score
}
func NewGame() *Tennis {
	return &Tennis{}
}
