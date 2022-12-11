package main

import (
	"fmt"
	"strconv"

	"github.com/manto89/aoc2022/common"
)

type Day08 struct {
	inputPath string
}

func (d *Day08) executePart1() string {
	lines, err := common.ReadAllLines(d.inputPath)

	if err != nil {
		return fmt.Sprintf("open file error: %v", err)
	}
	lastMaxCol := []int{}
	lastMaxColPos := []int{}
	cols := make([][]int, len(lines[0])-2)
	for i := 0; i < len(lines[0])-2; i++ {
		cols[i] = make([]int, len(lines))
	}
	visibles := make([][]bool, len(lines))
	for i := 0; i<len(lines); i++{
		visibles[i] = make([]bool, len(lines[0]))
		if i == 0 || i == len(lines)-1{
			for j := 0; j < len(lines[0]); j++ {
				//the first and the last row are visible
				visibles[i][j] = true
			}
		}
		visibles[i][0] = true
		visibles[i][len(lines[0])-1] = true

	}


	for i, line := range lines {
		lastMax, err := strconv.Atoi(string(line[0]))
		lastMaxPos := 0
		if err != nil {
			return fmt.Sprintf("cannot parse int at position 0 of line %d", i)
		}
		heights := []int{lastMax}
		//we know the first and last are already visible, skip
		for j := 1; j < len(line)-1; j++ {
			height, err := strconv.Atoi(string(line[j]))
			if err != nil {
				return fmt.Sprintf("cannot parse int at position %d of line %d", j, i)
			}
			//at the beginning lastMaxCol is empty
			if i == 0 {
				_, err := strconv.Atoi(string(line[len(line)-1]))
				if err != nil {
					return fmt.Sprintf("cannot parse int at position %d of line %d", len(line)-1, i)
				}

				
				lastMaxCol = append(lastMaxCol, height)
				lastMaxColPos = append(lastMaxColPos, 0)
			} else if lastMaxCol[j-1] < height {
				lastMaxCol[j-1] = height
				lastMaxColPos[j-1] = i
				//from top
				visibles[i][j] = true
			}

			if j > 0 && j < len(line)-1 {
				cols[j-1][i] = height
			}

			heights = append(heights, height)
			// we don't check the first and last line, they are all visible
			if i > 0 && i < len(lines)-1 && height > lastMax {
				//from left
				visibles[i][j] = true
				lastMax = height
				lastMaxPos = j
			}
		}
		lastPos := len(line) - 1
		lastMax, err = strconv.Atoi(string(line[lastPos]))
		if err != nil {
			return fmt.Sprintf("cannot parse int at position %d of line %d", lastPos, i)
		}
		heights = append(heights, lastMax)
		if i > 0 && i < len(lines)-1 {
			//we add the first and the last tree
			for j := lastPos - 1; j > lastMaxPos; j-- {
				if heights[j] > lastMax {
					//from right
					visibles[i][j] = true
					lastMax = heights[j]
				}
			}

		}

	}
	for i, col := range cols {
		lastMax := col[len(col)-1]
		for j := len(col) - 2; j > lastMaxColPos[i]; j-- {
			if col[j] > lastMax {
				//from bottom
				visibles[j][i+1] = true
				lastMax = col[j]
			}

		}
	}
	totalVisibles := 0
	for i := 0; i < len(visibles); i++{
		for j := 0; j < len(visibles[i]); j++{
			if visibles[i][j]{
				totalVisibles++
			}
		}
	}
	return fmt.Sprintf("number of visible trees: %d", totalVisibles)
}

func (d *Day08) executePart2() string {
	return "Not implemented"
}
