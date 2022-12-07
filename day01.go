package main

import (
	"fmt"
	"sort"
	"strconv"
	"github.com/manto89/aoc2022/common"
)

type Day01 struct {
	inputPath string
}

func (d *Day01) executePart1() string {

	lines, err := common.ReadAllLines(d.inputPath)

	if err != nil {
        return fmt.Sprintf("open file error: %v", err)
	}
	var caloriesList []int
	calories := 0
    for _, line := range(lines){
		if len(line) < 1 {
			caloriesList = append(caloriesList, calories)
			calories = 0
		}
		newCalories, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		calories += newCalories
    }
	sort.Ints(caloriesList)
	return strconv.Itoa(caloriesList[len(caloriesList)-1])

}

func (d *Day01) executePart2() string {
	return "Not implemented"
}