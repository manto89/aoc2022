package main

import (
	"errors"
	"fmt"

	"github.com/manto89/aoc2022/common"
)

type Day05 struct {
	inputPath string
}

func (d *Day05) executePart1() string{
	lines, err := common.ReadAllLines(d.inputPath)
	if err != nil {
        return fmt.Sprintf("open file error: %v", err)
	}
	indexDivider, err := getDividerIndex(lines)

	if err != nil {
		return err.Error()
	}
	//there is an additional line with stacks indexes, ignore it
	stacks, err := getStacks(lines[:indexDivider-1])
	if err != nil {
		return err.Error()
	}
	for i, line := range(lines[indexDivider+1:]){
		stacks, err = moveStacks(stacks, line)
		if err != nil {
			return fmt.Sprintf("Error while reading line %d\n%v", i, err)
		}
	}
	message := ""
	for _, stack := range(stacks){
		message += string(stack[len(stack)-1])
	}
	return fmt.Sprintf("Message is %s", message)
}

func moveStacks(stacks []string, instruction string) ([]string, error) {
	//move 3 from 9 to 7
	if len(instruction) < 1 {
		return nil, errors.New("instruction is too short")
	}
	if len(stacks) < 1{
		return nil, errors.New("there should be at least one stack")
	}
	var quantity, origin, destination int
	fmt.Sscanf(instruction, "move %d from %d to %d", &quantity, &origin, &destination)
	//arrays are 0 based
	origin -= 1
	destination -= 1
	for i := 0; i < quantity; i++ {
		crate := string(stacks[origin][len(stacks[origin])-1])
		stacks[origin] = stacks[origin][:len(stacks[origin])-1]
		stacks[destination] = stacks[destination] + string(crate)
	}
	return stacks, nil
}
func moveStacksV2(stacks []string, instruction string) ([]string, error) {
	//move 3 from 9 to 7
	if len(instruction) < 1 {
		return nil, errors.New("instruction is too short")
	}
	if len(stacks) < 1{
		return nil, errors.New("there should be at least one stack")
	}
	var quantity, origin, destination int
	fmt.Sscanf(instruction, "move %d from %d to %d", &quantity, &origin, &destination)
	//arrays are 0 based
	origin -= 1
	destination -= 1
	crates := string(stacks[origin][len(stacks[origin])-quantity:])
	stacks[origin] = stacks[origin][:len(stacks[origin])-quantity]
	stacks[destination] = stacks[destination] + crates
	return stacks, nil
}

func getDividerIndex(lines []string) (int, error) {
	indexDivider := 0
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			indexDivider = i
			break
		}
	}
	if indexDivider < 1{
		return 0, errors.New("unable to find divider between crates and instructions")
	}
	return indexDivider, nil
}

func getStacks(lines []string) ([]string, error){
	//[R] [H] [Z] [M] [T] [M] [T] [Q] [W]
	exampleStack := "[R] "
	numStacks := len(lines[0])/len(exampleStack)+1
	if numStacks < 1{
		return nil, errors.New("the first line is too short")
	}
	stacks := make([]string, len(lines[0])/len(exampleStack)+1)
	for i, line := range(lines){
		if len(line) / len(exampleStack) +1 != numStacks{
			return nil, fmt.Errorf("the stacks on line %d are not the same" + 
			"as the ones on the first line (%d)", i, numStacks)
		}
		stackCounter := 0
		for j := 1; j < len(line); j+=len(exampleStack) {
			crate := string(line[j])
			if crate != " "{
				stacks[stackCounter] += crate
			}
			stackCounter += 1
		}
	}
	for i := 0; i < len(stacks); i++ {
		stacks[i] = common.ReverseString(stacks[i])
	}
	return stacks, nil

}



func (d *Day05) executePart2() string{
	lines, err := common.ReadAllLines(d.inputPath)
	if err != nil {
        return fmt.Sprintf("open file error: %v", err)
	}
	indexDivider, err := getDividerIndex(lines)

	if err != nil {
		return err.Error()
	}
	//there is an additional line with stacks indexes, ignore it
	stacks, err := getStacks(lines[:indexDivider-1])
	if err != nil {
		return err.Error()
	}
	for i, line := range(lines[indexDivider+1:]){
		stacks, err = moveStacksV2(stacks, line)
		if err != nil {
			return fmt.Sprintf("Error while reading line %d\n%v", i, err)
		}
	}
	message := ""
	for _, stack := range(stacks){
		message += string(stack[len(stack)-1])
	}
	return fmt.Sprintf("Message is %s", message)
}