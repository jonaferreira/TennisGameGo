package main

import (
	"fmt"
	"testing"
)

type AddResult struct {
	x        int
	y        int
	expected int
}

var AddResults = []AddResult{
	{1, 1, 2},
}

func TestScoring(t *testing.T) {
	game := Game{Player{name: "DelPotro"},
		Player{name: "Djokovic"}}

	scoring(&game.player1)
	fmt.Println(game)
	scoring(&game.player1)
	fmt.Println(game)
	scoring(&game.player1)
	fmt.Println(game)

}

/*func TestAdd(t *testing.T) {
	for _, test := range AddResults {
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Fatal("Expected Result Not Given")
		}
	}
}*/
