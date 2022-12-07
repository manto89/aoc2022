package common

import (
	"fmt"
)

type RockPaperScissor int

const (
	Rock RockPaperScissor = iota
	Paper
	Scissor
)

func (d RockPaperScissor) String() string {
	return [...]string{"Rock", "Paper", "Scissor", }[d]
}

type RockPaperScissorChoice struct {
	choice RockPaperScissor
	Score int
	inputFirstPlayer string
	inputSecondPlayer string
	winsAgainst RockPaperScissor
	loosesAgainst RockPaperScissor
}

type RockPaperScissorReader struct {
	allChoices []RockPaperScissorChoice 
}

func (reader *RockPaperScissorReader) GetFirstChoice(input string) (*RockPaperScissorChoice, error){
	for _, choice := range(reader.allChoices){
		if choice.inputFirstPlayer == input {
			return &choice, nil
		}
	}
	return nil, fmt.Errorf("unable to create choice from input %s", input )
}
func (reader *RockPaperScissorReader) GetSecondChoice(input string) (*RockPaperScissorChoice, error){
	for _, choice := range(reader.allChoices){
		if choice.inputSecondPlayer == input {
			return &choice, nil
		}
	}
	return nil, fmt.Errorf("unable to create choice from input %s", input )
}

func MakeReader() *RockPaperScissorReader{
	ret := RockPaperScissorReader{
		allChoices : []RockPaperScissorChoice{
	GetRock(),
	GetPaper(), 
	GetScissor(),
		},
	}
	return &ret
}

func GetRock() RockPaperScissorChoice {
	return RockPaperScissorChoice{
		choice: Rock, 
		Score: 1, 
		inputFirstPlayer: "A",
		inputSecondPlayer: "X",
		winsAgainst: Scissor,
		loosesAgainst: Paper,
	}
}
func GetPaper() RockPaperScissorChoice {
	return RockPaperScissorChoice{
		choice: Paper,
		Score: 2,
		inputFirstPlayer: "B",
		inputSecondPlayer: "Y",
		winsAgainst: Rock,
		loosesAgainst: Scissor,
	}
}
func GetScissor() RockPaperScissorChoice {
	return RockPaperScissorChoice{
		choice: Scissor,
		Score: 3,
		inputFirstPlayer: "C",
		inputSecondPlayer: "Z",
		winsAgainst: Paper,
		loosesAgainst: Rock,
	}
}

// this method will return 0 if second player (me) looses, 3 if tie, 6 if second player (me) wins
func BattleScore(firstPlayer RockPaperScissorChoice, me RockPaperScissorChoice) int {
	ret := 0
	if firstPlayer.choice == me.choice{
		ret = 3
	} else if me.winsAgainst == firstPlayer.choice{
		ret = 6
	} 
	return ret
}