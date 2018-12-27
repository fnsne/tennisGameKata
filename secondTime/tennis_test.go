package tennis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTennisGame(t *testing.T) {
	tennisGame := getTennisGame()
	assert.NotNil(t, tennisGame)
}
func TestLoveAll(t *testing.T) {
	tennisGame := getTennisGame()
	tennisGame.setPlayerAScore(0)
	tennisGame.setPlayerAScore(0)
	assert.Equal(t, "Love All", tennisGame.GetScore())
}
func TestFifteenLove(t *testing.T) {
	tennisGame := getTennisGame()
	tennisGame.setPlayerAScore(1)
	tennisGame.setPlayerBScore(0)
	assert.Equal(t, "Fifteen Love", tennisGame.GetScore())
}
func TestThirtyLove(t *testing.T) {
	tennisGame := getTennisGame()
	tennisGame.setPlayerAScore(2)
	tennisGame.setPlayerBScore(0)
	assert.Equal(t, "Thirty Love", tennisGame.GetScore())
}
func TestFortyLove(t *testing.T) {
	tennisGame := getTennisGame()
	tennisGame.setPlayerAScore(3)
	tennisGame.setPlayerBScore(0)
	assert.Equal(t, "Forty Love", tennisGame.GetScore())
}

func TestLoveFifteen(t *testing.T) {
	tennisGame := getTennisGame()
	tennisGame.setPlayerAScore(0)
	tennisGame.setPlayerBScore(1)
	assert.Equal(t, "Love Fifteen", tennisGame.GetScore())
}
func TestLoveThirty(t *testing.T) {
	tennisGame := getTennisGame()
	tennisGame.setPlayerAScore(0)
	tennisGame.setPlayerBScore(2)
	assert.Equal(t, "Love Thirty", tennisGame.GetScore())
}
func TestLoveForty(t *testing.T) {
	tennisGame := getTennisGame()
	tennisGame.setPlayerAScore(0)
	tennisGame.setPlayerBScore(3)
	assert.Equal(t, "Love Forty", tennisGame.GetScore())
}
func TestFifteenAll(t *testing.T) {
	tennisGame := getTennisGame()
	tennisGame.setPlayerAScore(1)
	tennisGame.setPlayerBScore(1)
	assert.Equal(t, "Fifteen All", tennisGame.GetScore())
}
func TestThirtyAll(t *testing.T) {
	tennisGame := getTennisGame()
	tennisGame.setPlayerAScore(2)
	tennisGame.setPlayerBScore(2)
	assert.Equal(t, "Thirty All", tennisGame.GetScore())
}
func TestDeuce3_3(t *testing.T) {
	tennisGame := getTennisGame()
	tennisGame.setPlayerAScore(3)
	tennisGame.setPlayerBScore(3)
	assert.Equal(t, "Deuce", tennisGame.GetScore())
}
func TestDeuce4_4(t *testing.T) {
	tennisGame := getTennisGame()
	tennisGame.setPlayerAScore(4)
	tennisGame.setPlayerBScore(4)
	assert.Equal(t, "Deuce", tennisGame.GetScore())
}

func Test4_3(t *testing.T) {
	tennisGame := getTennisGame()
	tennisGame.setPlayerAScore(4)
	tennisGame.setPlayerBScore(3)
	assert.Equal(t, "Player A Adv", tennisGame.GetScore())
}
func Test5_3(t *testing.T) {
	tennisGame := getTennisGame()
	tennisGame.setPlayerAScore(5)
	tennisGame.setPlayerBScore(3)
	assert.Equal(t, "Player A Win", tennisGame.GetScore())
}
func Test3_4(t *testing.T) {
	tennistGame := getTennisGame()
	tennistGame.setPlayerAScore(3)
	tennistGame.setPlayerBScore(4)
	assert.Equal(t, "Player B Adv", tennistGame.GetScore())
}
func Test3_5(t *testing.T) {
	tennisGame := getTennisGame()
	tennisGame.setPlayerAScore(3)
	tennisGame.setPlayerBScore(5)
	assert.Equal(t, "Player B Win", tennisGame.GetScore())
}
