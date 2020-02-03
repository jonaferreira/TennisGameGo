package main

import (
	"fmt"
	"math/rand"
	"time"
)

type statesScoring struct {
	Point           int
	MatchPoint      int
	AdvantangePoint int
	WinnerPoint     int
	DeusePoint      int
}

// EnumStatesScoring for public use
var EnumStatesScoring = &statesScoring{
	Point:           0,
	MatchPoint:      1,
	AdvantangePoint: 2,
	WinnerPoint:     3,
	DeusePoint:      4,
}

// Points is ...
var Points = [4]int{0, 15, 30, 40}

// Player is ...
type Player struct {
	name      string
	score     int
	winner    bool
	advantage bool
	service   bool
}

// PlayerInterface is ...
type PlayerInterface interface {
	getName() string
	getScore() int
	hasAdvantage() bool
	hasService() bool
	winScore()
	setScore(score int)
	setAdvantage(advantage bool)
	setWinner(winner bool)
	setService(service bool)
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

func (player *Player) hasService() bool {
	return player.service
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

func (player *Player) setService(service bool) {
	player.service = service
}

func (player *Player) triesReturnBall() bool {
	fmt.Println(player.getName(), "is going to the ball... and...")

	if rand.Float64() < 0.5 {
		fmt.Println(player.getName(), "can't return the ball")
		return false
	}

	fmt.Println(player.getName(), "had returned the ball")
	return true
}

// Game is ....
type Game struct {
	player1 Player
	player2 Player
}

func (game *Game) isThereWinner() bool {
	return !(game.player1.winner == true || game.player2.winner == true)
}

func (game *Game) calculateService(turn int) int {

	if game.player1.hasService() {
		fmt.Println(game.player1.getName(), " is serving...")
		turn = 0
		game.player1.setService(false)
	} else {
		turn = turn % 2
		turn++
	}

	return turn

}

func scoring(playerWin PlayerInterface, playerLose PlayerInterface) int {
	state := EnumStatesScoring.Point

	if playerWin.getScore() < 2 {
		playerWin.winScore()
	} else if playerWin.hasAdvantage() == false && playerLose.hasAdvantage() == false {
		if playerLose.getScore() < 3 {
			state = EnumStatesScoring.MatchPoint
		} else {
			state = EnumStatesScoring.AdvantangePoint
		}
		playerWin.setScore(3)
		playerWin.setAdvantage(true)
	} else if playerWin.hasAdvantage() == true && playerLose.hasAdvantage() == false {
		state = EnumStatesScoring.WinnerPoint
		playerWin.setWinner(true)
	} else if playerWin.hasAdvantage() == false && playerLose.hasAdvantage() == true {
		state = EnumStatesScoring.DeusePoint
		playerWin.setScore(3)
		playerLose.setScore(3)
		playerLose.setAdvantage(false)
	}

	return state
}

func printJudgeSays(state int, playerName string) {
	fmt.Println("Gool!!.. Point to", playerName)

	switch state {
	case EnumStatesScoring.MatchPoint:
		fmt.Println("Math Point to", playerName)
	case EnumStatesScoring.AdvantangePoint:
		fmt.Println("Advantage to", playerName, ".. Match Point!!!")
	case EnumStatesScoring.WinnerPoint:
		fmt.Println("Winner is", playerName)
	case EnumStatesScoring.DeusePoint:
		fmt.Println("Deuse!!!")
	}
}

func (game *Game) playing() bool {
	continues := true

	game.player1.setService(true)
	turn := 1

	for continues {

		turn = game.calculateService(turn)

		if turn == 1 {
			if !game.player1.triesReturnBall() {
				state := scoring(&game.player2, &game.player1)
				printJudgeSays(state, game.player2.getName())
				fmt.Println(game.player1, " | ", game.player2)
				game.player1.setService(true)
			}
		} else {
			if !game.player2.triesReturnBall() {
				state := scoring(&game.player1, &game.player2)
				printJudgeSays(state, game.player1.getName())
				fmt.Println(game.player1, " | ", game.player2)
				game.player1.setService(true)
			}
		}
		continues = game.isThereWinner()
	}
	return continues
}

func (game *Game) round() bool {
	fmt.Println("Start new round!!...")
	return game.playing()
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
