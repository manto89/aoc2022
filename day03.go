package main

import (
	"fmt"
	"github.com/manto89/aoc2022/common"
)

type Day03 struct{
	inputPath string
}

func (d *Day03) executePart1() string{
	lines, err := common.ReadAllLines(d.inputPath)
	if err != nil {
        return fmt.Sprintf("open file error: %v", err)
	}
	priorityScore := 0
	for i, line := range(lines){
		var items [52]bool
		compartment1 := line[:(len(line)/2)]
		compartment2 := line[(len(line)/2):]

		for _, char := range(compartment1){
			priority, err := GetPriority(char)
			if err != nil {
				return fmt.Sprintf("Unable to decode character at line %d", i)
			}
			items[priority-1] = true
		}
		for _, char := range(compartment2){
			priority, err := GetPriority(char)
			if err != nil {
				return fmt.Sprintf("Unable to decode character at line %d", i)
			}
			if items[priority-1] {
				priorityScore += priority
				break

			}

		}

	}
	return fmt.Sprintf("Priority score is %d", priorityScore)

}

func GetPriority(c rune) (int, error) {
	priority := 0
	ascii := int(c)
	if ascii > 96 && ascii < 123 {
		priority = ascii - 96
	} else if ascii > 64 && ascii < 91{
		priority = ascii - 38
	} else {
		return 0, fmt.Errorf("unable to decode character %c", c)
	}
	return priority, nil
}

func (d *Day03) executePart2() string{
	lines, err := common.ReadAllLines(d.inputPath)
	if err != nil {
        return fmt.Sprintf("open file error: %v", err)
	}
	priorityScore := 0
	var groupItems [][52]bool
	for i, line := range(lines){
		var items [52]bool

		for _, char := range(line){
			priority, err := GetPriority(char)
			if err != nil {
				return fmt.Sprintf("Unable to decode character at line %d", i)
			}
			items[priority-1] = true
		}
		groupItems = append(groupItems, items)
		if i % 3 == 2 {
			for i, _ := range(items){
				if groupItems[0][i] && groupItems[1][i] && groupItems[2][i]{
					//priority starts from 1, arrays start from 0
					priorityScore += i+1
					break
				}
			}
			groupItems = [][52]bool{}
		}

	}
	return fmt.Sprintf("Priority score is %d", priorityScore)
}
