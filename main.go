package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Points = [4]int{0, 15, 30, 40}

type Player struct {
	name       string
	score      int
	matchPoint bool
}

type Game struct {
	player1 Player
	player2 Player
}

func scoring(player *Player) {
	defer player.isMatchPoint()
	player.score++
}

func (player *Player) isMatchPoint() {
	if player.score == 3 {
		player.matchPoint = true
		fmt.Println("Match Point for: ", player)
	}
}

func (game *Game) wonPoint() {
	if rand.Float64() < 0.5 {
		scoring(&game.player1)
	} else {
		scoring(&game.player2)
	}
}

func (game *Game) isThereDeuse() {
	if game.player1.matchPoint && game.player2.matchPoint {
		game.player1.matchPoint = false
		game.player2.matchPoint = false
		fmt.Println("Deuseee!!")
	} else {
		fmt.Println("There isn't deuse")
	}
}

func (game *Game) isThereWinner() {

}

func (game *Game) round() bool {
	game.wonPoint()
	game.isThereDeuse()
	game.isThereWinner()
	return false
}

func main() {
	rand.Seed(time.Now().Unix())

	game := Game{Player{name: "DelPotro", score: 2},
		Player{name: "Djokovic", score: 3}}

	continues := true
	for continues {
		continues = game.round()
		fmt.Println(game)
	}
}

func (player Player) String() string {
	return fmt.Sprint("name: ", player.name,
		" - score: ", player.score,
		" - isYourMatchPoint: ", player.matchPoint)
}
