package main

import (
	"fmt"

	"github.com/manto89/aoc2022/common"
)
type Day06 struct{
	inputPath string
}
func (d *Day06) executePart1() string {
	lines, err := common.ReadAllLines(d.inputPath)
	if err != nil {
        return fmt.Sprintf("open file error: %v", err)
	}
	line := lines[0]
	marker := line[:4]
	markerIndex := 4
	for i := 4; i < len(line); i++ {
		if allCharactersAreDifferent(marker){
			markerIndex = i
			break;
		}
		marker = marker[1:] + string(line[i])
	}
	return fmt.Sprintf("Marker found at %d", markerIndex)

}

func allCharactersAreDifferent(marker string) bool{
	ret := true
	for i := 0; i < len(marker); i++ {
		//compare the character with all the characters after (not the ones before, it's useless)
		for j := i+1; j < len(marker); j++{
			if marker[i] == marker[j]{
				ret = false
				break
			}
		}
		if !ret {
			break
		}
	}

	return ret
}
func (d *Day06) executePart2() string {
	return "Not implemented"
}