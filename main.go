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

// Game is ....
type Game struct {
	player1 Player
	player2 Player
}

func (game *Game) isThereWinner() bool {
	return !(game.player1.winner == true || game.player2.winner == true)
}

func scoring(playerWin *Player, playerLose *Player) {
	fmt.Println("Gool!!.. Point to", playerWin.name)
	if playerWin.score < 2 {
		playerWin.score++
	} else if playerWin.advantage == false && playerLose.advantage == false {
		if playerLose.score < 3 {
			fmt.Println("Math Point to", playerWin.name)
		} else {
			fmt.Println("Advantage to", playerWin.name, ".. Match Point!!!")
		}
		playerWin.score = 3
		playerWin.advantage = true
	} else if playerWin.advantage == true && playerLose.advantage == false {
		fmt.Println("Winner is", playerWin.name)
		playerWin.winner = true
	} else if playerWin.advantage == false && playerLose.advantage == true {
		fmt.Println("Deuse!!!")
		playerWin.score = 3
		playerLose.score = 3
		playerLose.advantage = false
	}
}

func (game *Game) wonPoint() {
	if rand.Float64() < 0.5 {
		scoring(&game.player1, &game.player2)
	} else {
		scoring(&game.player2, &game.player1)
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
