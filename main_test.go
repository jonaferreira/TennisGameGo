package main

type scoringResult struct {
	playerIn1  Player
	playerIn2  Player
	playerOut1 Player
	playerOut2 Player
}

var scoringResults = []scoringResult{
	{Player{score: 0}, Player{score: 0}, Player{score: 1}, Player{score: 0}},
	{Player{score: 1}, Player{score: 0}, Player{score: 2}, Player{score: 0}},
	{Player{score: 2}, Player{score: 0}, Player{score: 3}, Player{score: 0}},
	{Player{score: 3}, Player{score: 0}, Player{score: 4}, Player{score: 0}},
	{Player{score: 3}, Player{score: 0}, Player{score: 3}, Player{score: 1}},
	{Player{score: 3}, Player{score: 1}, Player{score: 3}, Player{score: 2}},
	{Player{score: 3}, Player{score: 2}, Player{score: 3}, Player{score: 3}},
	{Player{score: 3}, Player{score: 3}, Player{score: 3, advantage: true}, Player{score: 3, advantage: false}},
}

/*
func TestScoring(t *testing.T) {

}*/

/*func TestAdd(t *testing.T) {
	for _, test := range AddResults {
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Fatal("Expected Result Not Given")
		}
	}
}*/
