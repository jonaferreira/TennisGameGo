package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Points is ...
var Points = [4]int{0, 15, 30, 40}

// Player is ...
type Player struct {
	name      string
	score     int
	winner    bool
	advantage bool
}

type PlayerInterface interface {
	getName() string
	getScore() int
	hasAdvantage() bool
	winScore()
	setScore(score int)
	setAdvantage(advantage bool)
	setWinner(winner bool)
}

func (player *Player) getName() string {
	return player.name
}

func (player *Player) getScore() int {
	return player.score
}

func (player *Player) hasAdvantage() bool {
	return player.advantage
}

func (player *Player) winScore() {
	player.score++
}

func (player *Player) setScore(score int) {
	player.score = score
}

func (player *Player) setAdvantage(advantage bool) {
	player.advantage = advantage
}

func (player *Player) setWinner(winner bool) {
	player.winner = winner
}

// Game is ....
type Game struct {
	player1 Player
	player2 Player
}

func (game *Game) isThereWinner() bool {
	return !(game.player1.winner == true || game.player2.winner == true)
}

func scoring(playerWin PlayerInterface, playerLose PlayerInterface) {
	fmt.Println("Gool!!.. Point to", playerWin.getName())
	if playerWin.getScore() < 2 {
		playerWin.winScore()
	} else if playerWin.hasAdvantage() == false && playerLose.hasAdvantage() == false {
		if playerLose.getScore() < 3 {
			fmt.Println("Math Point to", playerWin.getName())
		} else {
			fmt.Println("Advantage to", playerWin.getName(), ".. Match Point!!!")
		}
		playerWin.setScore(3)
		playerWin.setAdvantage(true)
	} else if playerWin.hasAdvantage() == true && playerLose.hasAdvantage() == false {
		fmt.Println("Winner is", playerWin.getName())
		playerWin.setWinner(true)
	} else if playerWin.hasAdvantage() == false && playerLose.hasAdvantage() == true {
		fmt.Println("Deuse!!!")
		playerWin.setScore(3)
		playerLose.setScore(3)
		playerLose.setAdvantage(false)
	}
}

func (game *Game) wonPoint() {
	if rand.Float64() < 0.5 {
		scoring(&game.player1, &game.player2)
	} else {
		scoring(&game.player2, &game.player1)
	}
}

func (game *Game) playGame() {
	continues := true
	turn := 1
	for continues {
		turn = turn % 2
		turn++
		//game.player1
	}
}

func (game *Game) round() bool {
	fmt.Println(game.player1, " | ", game.player2)
	game.wonPoint()
	return game.isThereWinner()
}

func main() {
	rand.Seed(time.Now().Unix())

	game := Game{Player{name: "DelPotro"},
		Player{name: "Djokovic"}}

	fmt.Println("Really!!... Go!!")

	continues := true
	for continues {
		continues = game.round()
	}
}

func (player Player) String() string {
	return fmt.Sprint("name: ", player.name,
		" - score: ", Points[player.score],
		" - hasAdvantage: ", player.advantage,
		" - isWinner: ", player.winner)
}
