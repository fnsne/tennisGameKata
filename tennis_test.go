package tennis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTennissGame(t *testing.T) {
	tenissGame, err := NewTennisGame("", "")
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.NotNil(t, tenissGame)
}

func TestLoveAll(t *testing.T) {
	tenissGame, _ := NewTennisGame("", "")
	score := tenissGame.GetScore()
	assert.NotNil(t, tenissGame)
	assert.Equal(t, "Love All", score)
}

func TestFifteenLove(t *testing.T) {
	tennisGame, _ := NewTennisGame("", "")
	assert.NotNil(t, tennisGame)
	tennisGame.AScore = 1
	score := tennisGame.GetScore()
	assert.Equal(t, "Fifteen Love", score)
}

func TestThirtyLove(t *testing.T) {
	tennisGame, _ := NewTennisGame("", "")
	assert.NotNil(t, tennisGame)
	tennisGame.AScore = 2
	score := tennisGame.GetScore()
	assert.Equal(t, "Thirty Love", score)
}

func TestFortyLove(t *testing.T) {
	tennisGame, _ := NewTennisGame("", "")
	assert.NotNil(t, tennisGame)
	tennisGame.AScore = 3
	score := tennisGame.GetScore()
	assert.Equal(t, "Forty Love", score)
}

func TestLoveFifteen(t *testing.T) {
	tennisGame, _ := NewTennisGame("", "")
	assert.NotNil(t, tennisGame)
	tennisGame.BScore = 1
	assert.Equal(t, "Love Fifteen", tennisGame.GetScore())
}
func TestLoveThirty(t *testing.T) {
	tennisGame, _ := NewTennisGame("", "")
	assert.NotNil(t, tennisGame)
	tennisGame.BScore = 2
	assert.Equal(t, "Love Thirty", tennisGame.GetScore())
}
func TestLoveForty(t *testing.T) {
	tennisGame, _ := NewTennisGame("", "")
	assert.NotNil(t, tennisGame)
	tennisGame.BScore = 3
	assert.Equal(t, "Love Forty", tennisGame.GetScore())
}

func TestFifteenAll(t *testing.T) {
	tennisGame, _ := NewTennisGame("", "")
	assert.NotNil(t, tennisGame)
	tennisGame.AScore = 1
	tennisGame.BScore = 1
	assert.Equal(t, "Fifteen All", tennisGame.GetScore())
}

func TestThirtyAll(t *testing.T) {
	tennisGame, _ := NewTennisGame("", "")
	assert.NotNil(t, tennisGame)
	tennisGame.AScore = 2
	tennisGame.BScore = 2
	assert.Equal(t, "Thirty All", tennisGame.GetScore())
}
func TestDeuce3_3(t *testing.T) {
	tennisGame, _ := NewTennisGame("", "")
	assert.NotNil(t, tennisGame)
	tennisGame.AScore = 3
	tennisGame.BScore = 3
	assert.Equal(t, "Deuce", tennisGame.GetScore())
}
func TestDeuce4_4(t *testing.T) {
	tennisGame, _ := NewTennisGame("", "")
	assert.NotNil(t, tennisGame)
	tennisGame.AScore = 4
	tennisGame.BScore = 4
	assert.Equal(t, "Deuce", tennisGame.GetScore())
}
func TestJackAdv(t *testing.T) {
	tenissGame, _ := NewTennisGame("Jack", "Petter")
	assert.NotNil(t, tenissGame)
	tenissGame.AScore = 4
	tenissGame.BScore = 3
	assert.Equal(t, "Jack Adv", tenissGame.GetScore())
}
func TestJackWin(t *testing.T) {
	tennisGame, _ := NewTennisGame("Jack", "Petter")
	assert.NotNil(t, tennisGame)
	tennisGame.AScore = 5
	tennisGame.BScore = 3
	assert.Equal(t, "Jack Win", tennisGame.GetScore())
}
func TestPetterAdv(t *testing.T) {
	tennisGame, _ := NewTennisGame("Jack", "Petter")
	assert.NotNil(t, tennisGame)
	tennisGame.AScore = 3
	tennisGame.BScore = 4
	assert.Equal(t, "Petter Adv", tennisGame.GetScore())
}
func TestPetterWin(t *testing.T) {
	tennisGame, _ := NewTennisGame("Jack", "Petter")
	assert.NotNil(t, tennisGame)
	tennisGame.AScore = 3
	tennisGame.BScore = 5
	assert.Equal(t, "Petter Win", tennisGame.GetScore())
}
