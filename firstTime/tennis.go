package tennis

type TennisGame struct {
	AScore int
	BScore int
	AName  string
	BName  string
}

var SCORE map[int]string = getScoreName()

func NewTennisGame(a, b string) (*TennisGame, error) {
	return &TennisGame{
		AName:  a,
		BName:  b,
		AScore: 0,
		BScore: 0,
	}, nil
}

func (t TennisGame) GetScore() string {
	//分數相同
	if t.sameScore() {
		return t.sameScoreResponse()
	} else {
		if t.scoreEnough() {
			if abs(t.AScore-t.BScore) == 2 {
				return t.superiorName() + " Win"
			} else if abs(t.AScore-t.BScore) == 1 {
				return t.superiorName() + " Adv"
			}
		}
		return SCORE[t.AScore] + " " + SCORE[t.BScore]
	}

}
func (t TennisGame) scoreEnough() bool {
	if t.AScore >= 3 || t.BScore >= 3 {
		return true
	} else {
		return false
	}
}
func (t TennisGame) superiorName() string {
	if t.AScore > t.BScore {
		return t.AName
	} else {
		return t.BName
	}
}
func (t TennisGame) sameScoreResponse() string {
	if t.AScore >= 3 {
		return "Deuce"
	} else {
		return SCORE[t.AScore] + " All"
	}

}
func (t TennisGame) sameScore() bool {
	if t.AScore == t.BScore {
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
func abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}
