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

func convertToIntAndSumBetweenEmptyLines(strings []string) ([]int, error){
	var ret []int
	sum := 0
	for _, s := range(strings){
		if len(s) < 1 {
			ret = append(ret, sum)
			sum = 0
			continue;
		}
		i, err := strconv.Atoi(s)
		if (err != nil){
			return nil, err
		}
		sum += i
	}
	return ret, nil
}

func (d *Day01) executePart1() string {

	lines, err := common.ReadAllLines(d.inputPath)

	if err != nil {
        return fmt.Sprintf("open file error: %v", err)
	}

	ints, err := convertToIntAndSumBetweenEmptyLines(lines)

	if err != nil {
		return fmt.Sprintf("Unable to parse strings to ints %w", err)
	}

	sort.Ints(ints)
	return strconv.Itoa(ints[len(ints)-1])

}

func (d *Day01) executePart2() string {

	lines, err := common.ReadAllLines(d.inputPath)

	if err != nil {
        return fmt.Sprintf("open file error: %v", err)
	}

	ints, err := convertToIntAndSumBetweenEmptyLines(lines)

	if err != nil {
		return fmt.Sprintf("Unable to parse strings to ints %w", err)
	}

	sort.Ints(ints)
	sum := 0
	for _,i := range(ints[len(ints)-3:]){
		sum += i
	}
	
	return strconv.Itoa(sum)

}