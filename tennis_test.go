package Tennis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerTwoWin(t *testing.T) {
	game := NewTennisGame()
	game.setPlayerOneScore(3)
	game.setPlayerTwoScore(5)
	assert.Equal(t, "Player Two Win", game.GetScore())
}
func TestPlayerTwoAdv(t *testing.T) {
	game := NewTennisGame()
	game.setPlayerOneScore(3)
	game.setPlayerTwoScore(4)
	assert.Equal(t, "Player Two Adv", game.GetScore())
}
func TestPlayerOneWin(t *testing.T) {
	game := NewTennisGame()
	game.setPlayerOneScore(5)
	game.setPlayerTwoScore(3)
	assert.Equal(t, "Player One Win", game.GetScore())
}
func TestPlayerOneAdv(t *testing.T) {
	game := NewTennisGame()
	game.setPlayerOneScore(4)
	game.setPlayerTwoScore(3)
	assert.Equal(t, "Player One Adv", game.GetScore())
}
func TestDeuce_4_4(t *testing.T) {
	game := NewTennisGame()
	game.setPlayerOneScore(4)
	game.setPlayerTwoScore(4)
	assert.Equal(t, "Deuce", game.GetScore())
}
func TestDeuce_3_3(t *testing.T) {
	game := NewTennisGame()
	game.setPlayerOneScore(3)
	game.setPlayerTwoScore(3)
	assert.Equal(t, "Deuce", game.GetScore())
}
func TestThirtyAll(t *testing.T) {
	game := NewTennisGame()
	game.setPlayerOneScore(2)
	game.setPlayerTwoScore(2)
	assert.Equal(t, "Thirty All", game.GetScore())
}
func TestFifteenAll(t *testing.T) {
	game := NewTennisGame()
	game.setPlayerOneScore(1)
	game.setPlayerTwoScore(1)
	assert.Equal(t, "Fifteen All", game.GetScore())
}
func TestLoveForty(t *testing.T) {
	game := NewTennisGame()
	game.setPlayerOneScore(0)
	game.setPlayerTwoScore(3)
	assert.Equal(t, "Love Forty", game.GetScore())
}
func TestLoveThirty(t *testing.T) {
	game := NewTennisGame()
	game.setPlayerOneScore(0)
	game.setPlayerTwoScore(2)
	assert.Equal(t, "Love Thirty", game.GetScore())
}
func TestLoveFifteen(t *testing.T) {
	game := NewTennisGame()
	game.setPlayerOneScore(0)
	game.setPlayerTwoScore(1)
	assert.Equal(t, "Love Fifteen", game.GetScore())
}
func TestFortyLove(t *testing.T) {
	game := NewTennisGame()
	game.setPlayerOneScore(3)
	game.setPlayerTwoScore(0)
	assert.Equal(t, "Forty Love", game.GetScore())
}
func TestThirtyLove(t *testing.T) {
	game := NewTennisGame()
	game.setPlayerOneScore(2)
	game.setPlayerTwoScore(0)
	assert.Equal(t, "Thirty Love", game.GetScore())
}
func TestFifteenLove(t *testing.T) {
	game := NewTennisGame()
	game.setPlayerOneScore(1)
	game.setPlayerTwoScore(0)
	assert.Equal(t, "Fifteen Love", game.GetScore())
}
func TestLoveAll(t *testing.T) {
	game := NewTennisGame()
	game.setPlayerOneScore(0)
	game.setPlayerTwoScore(0)
	assert.Equal(t, "Love All", game.GetScore())
}
func TestGame(t *testing.T) {
	game := NewTennisGame()
	assert.NotNil(t, game)
}
