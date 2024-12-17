package services

import (
	"advent-of-code-2024/pkg/file"
	"advent-of-code-2024/pkg/matrix"
	"fmt"
	"strconv"
	"strings"
)

type IDay6Part2 interface {
	Solve() string
}

type Day6Part2 struct{}

func NewDay6Part2() IDay6Part2 {
	return &Day6Part2{}
}

func (obj *Day6Part2) Solve() string {
	fileString := file.ReadFileAsString()

	splitFileString := strings.Split(fileString, "\n")

	grid := make([][]string, len(splitFileString)) // (Y, X) cba transposing
	var startX int
	var startY int
	for index, line := range splitFileString {
		reversedIndex := len(splitFileString) - 1 - index

		grid[reversedIndex] = strings.Split(line, "")

		for lineIndex, _ := range grid[reversedIndex] {
			if grid[reversedIndex][lineIndex] == "^" {
				startX = lineIndex
				startY = reversedIndex
			}
		}
	}

	sum := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == "#" {
				continue
			}

			grid[i][j] = "#"

			position := [][]int{{startX}, {startY}}
			direction := [][]int{{0}, {1}}
			turn90 := [][]int{{0, 1}, {-1, 0}}

			hasLoop := false
			turns := map[string]string{}

			for true {

				// whats ahead
				aheadPosition := matrix.Add2d(position, direction)
				xAhead := aheadPosition[0][0]
				yAhead := aheadPosition[1][0]

				if xAhead >= len(grid[0]) || yAhead >= len(grid) || xAhead < 0 || yAhead < 0 {
					break
				}

				stringAhead := grid[yAhead][xAhead]

				if stringAhead == "#" {
					key := strconv.Itoa(position[0][0]) + strconv.Itoa(position[1][0]) + strconv.Itoa(direction[0][0]) + strconv.Itoa(direction[1][0])

					_, containsKey := turns[key]

					if containsKey {
						hasLoop = true
						break
					}

					turns[key] = ""

					direction = matrix.Multiply2d(turn90, direction)
					continue
				}

				position = matrix.Add2d(position, direction)
			}

			grid[i][j] = "."

			if hasLoop {
				sum++
			}
		}
	}

	return strconv.Itoa(sum)
}

func PrintGrid(grid [][]string) {
	for i := len(grid) - 1; i >= 0; i-- {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Print("\n")
	}
}
