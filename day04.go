package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/manto89/aoc2022/common"
)

type Day04 struct{
	inputPath string
}

func (d *Day04) executePart1() string{
	lines, err := common.ReadAllLines(d.inputPath)
	if err != nil {
        return fmt.Sprintf("open file error: %v", err)
	}
	sameSections := 0
	for i, line := range(lines){
		splittedLine := strings.Split(line, ",")
		if len(splittedLine) < 2 {
			return fmt.Sprintf("Unable to decode elves in line %d", i)
		}
		firstElf := strings.Split(splittedLine[0],"-")
		if len(firstElf) < 2 {
			return fmt.Sprintf("Unable to decode first elf in line %d", i)
		}
		secondElf := strings.Split(splittedLine[1],"-")
		if len(secondElf) < 2 {
			return fmt.Sprintf("Unable to decode second elf in line %d", i)
		}
		firstElfFirstSection, err := strconv.Atoi(firstElf[0])
		if err != nil {
			return fmt.Sprintf("Unable to decode sections in line %d", i)
		}
		firstElfSecondSection, err := strconv.Atoi(firstElf[1])
		if err != nil {
			return fmt.Sprintf("Unable to decode sections in line %d", i)
		}
		secondElfFirstSection, err := strconv.Atoi(secondElf[0])
		if err != nil {
			return fmt.Sprintf("Unable to decode sections in line %d", i)
		}
		secondElfSecondSection, err := strconv.Atoi(secondElf[1])
		if err != nil {
			return fmt.Sprintf("Unable to decode sections in line %d", i)
		}
		if (firstElfFirstSection >= secondElfFirstSection && firstElfSecondSection <= secondElfSecondSection) ||
		   (secondElfFirstSection >= firstElfFirstSection && secondElfSecondSection <= firstElfSecondSection) {
			sameSections += 1
		}

	}
	return fmt.Sprintf("Elves couple doing the same sections are %d", sameSections)
}
func (d *Day04) executePart2() string{
	return "Not implemented"

}