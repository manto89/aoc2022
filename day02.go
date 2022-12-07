package main

import (
	"fmt"
	"strings"

	"github.com/manto89/aoc2022/common"
)

type Day02 struct{
	inputPath string
}


func (d *Day02) executePart1() string{
	lines, err := common.ReadAllLines(d.inputPath)
	if err != nil {
        return fmt.Sprintf("open file error: %v", err)
	}
	reader := common.MakeReader()
	finalScore := 0
	for i, line := range(lines){
		splittedLine := strings.Split(line, " ")
		if len(splittedLine) < 2{
			return fmt.Sprintf("error splitting line %d", i)
		}
		firstChoice, err := reader.GetFirstChoice(splittedLine[0])
		if err != nil {
			return fmt.Sprintf("error reading first choice on line %d", i)
		}
		secondChoice, err := reader.GetSecondChoice(splittedLine[1])
		if err != nil {
			return fmt.Sprintf("error reading second choice on line %d", i)
		}

		battleScore := common.BattleScore(*firstChoice, *secondChoice)
		finalScore += battleScore + secondChoice.Score
		
	}
	return fmt.Sprintf("Final score was %d", finalScore)

}

func (d *Day02) executePart2() string{
	return "Not implemented"
}



