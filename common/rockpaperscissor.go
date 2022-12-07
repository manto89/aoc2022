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
	Choice RockPaperScissor
	Score int
	inputFirstPlayer string
	inputSecondPlayer string
	WinsAgainst RockPaperScissor
	LoosesAgainst RockPaperScissor
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
func GetFromEnum(enum RockPaperScissor) RockPaperScissorChoice{
	if enum == Rock{
		return GetRock()
	} else if enum == Paper{
		return GetPaper()
	} else{
		return GetScissor()
	}
}

func GetRock() RockPaperScissorChoice {
	return RockPaperScissorChoice{
		Choice: Rock, 
		Score: 1, 
		inputFirstPlayer: "A",
		inputSecondPlayer: "X",
		WinsAgainst: Scissor,
		LoosesAgainst: Paper,
	}
}
func GetPaper() RockPaperScissorChoice {
	return RockPaperScissorChoice{
		Choice: Paper,
		Score: 2,
		inputFirstPlayer: "B",
		inputSecondPlayer: "Y",
		WinsAgainst: Rock,
		LoosesAgainst: Scissor,
	}
}
func GetScissor() RockPaperScissorChoice {
	return RockPaperScissorChoice{
		Choice: Scissor,
		Score: 3,
		inputFirstPlayer: "C",
		inputSecondPlayer: "Z",
		WinsAgainst: Paper,
		LoosesAgainst: Rock,
	}
}

// this method will return 0 if second player (me) looses, 3 if tie, 6 if second player (me) wins
func BattleScore(firstPlayer RockPaperScissorChoice, me RockPaperScissorChoice) int {
	ret := 0
	if firstPlayer.Choice == me.Choice{
		ret = 3
	} else if me.WinsAgainst == firstPlayer.Choice{
		ret = 6
	} 
	return ret
}