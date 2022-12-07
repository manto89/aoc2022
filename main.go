package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1{
		print("missing day number. Add day as argument, eg. 'main.go 1' for day 1")
		return
	}
	dayString := args[0]
	dayNum, err := strconv.Atoi(dayString)
	if err != nil {
		print("the provided argument is not a number")
		return
	}
	if dayNum < 1 {
		print("the argument should be a positive number")
		return
	}

	days := map[int]IDay{
		1: &Day01{inputPath: "inputs/day01.txt"},
		2: &Day02{inputPath: "inputs/day02.txt"},
		3: &Day03{inputPath: "inputs/day03.txt"},
		4: &Day03{inputPath: "inputs/day04.txt"},
	}	
	dayToExecute, ok := days[dayNum]
	if !ok {
		print(fmt.Sprintf("day %d not found", dayNum))
		return
	}
	println(fmt.Sprintf("Executing day %d ", dayNum))
	println("Part 1:")
	println(dayToExecute.executePart1())
	println("Part 2:")
	println(dayToExecute.executePart2())
}
