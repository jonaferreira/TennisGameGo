package main

import (
	"testing"
)

func TestScoringSimpleOk(t *testing.T) {

	player1 := Player{}
	player2 := Player{}

	player1Expect := Player{score: 1}

	scoring(&player1, &player2)

	if player1.getScore() != player1Expect.getScore() {
		t.Fatal("Expected Result Not Given")
	}
}

func TestScoringMatchPoint(t *testing.T) {

	player1 := Player{score: 3}
	player2 := Player{score: 3}

	player1Expect := Player{advantage: true}

	scoring(&player1, &player2)

	if player1.hasAdvantage() != player1Expect.hasAdvantage() {
		t.Fatal("Expected Result Not Given")
	}
}
