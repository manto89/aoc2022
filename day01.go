package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Day01 struct {
	inputPath string
}

func (d *Day01) executePart1() string {

	f, err := os.OpenFile(d.inputPath, os.O_RDONLY, os.ModePerm)
    if err != nil {
        return fmt.Sprintf("open file error: %v", err)
    }
    defer f.Close()

    sc := bufio.NewScanner(f)
	var caloriesList []int
	calories := 0
    for sc.Scan() {
        line := sc.Text()  // GET the line string
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
    if err := sc.Err(); err != nil {
        return fmt.Sprintf("scan file error: %v", err)
    }
	sort.Ints(caloriesList)
	return strconv.Itoa(caloriesList[len(caloriesList)-1])

}

func (d *Day01) executePart2() string {
	return "Not implemented"
}