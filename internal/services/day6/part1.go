package services

import (
	"advent-of-code-2024/pkg/file"
	"advent-of-code-2024/pkg/matrix"
	"strconv"
	"strings"
)

type IDay6Part1 interface {
	Solve() string
}

type Day6Part1 struct{}

func NewDay6Part1() IDay6Part1 {
	return &Day6Part1{}
}

func (obj *Day6Part1) Solve() string {
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

	position := [][]int{{startX}, {startY}}
	direction := [][]int{{0}, {1}}
	turn90 := [][]int{{0, 1}, {-1, 0}}

	for true {
		// whats ahead
		aheadPosition := matrix.Add2d(position, direction)
		xAhead := aheadPosition[0][0]
		yAhead := aheadPosition[1][0]

		if xAhead >= len(grid[0]) || yAhead >= len(grid) || xAhead < 0 || yAhead < 0 {
			x := position[0][0]
			y := position[1][0]

			grid[y][x] = "X"

			break
		}

		stringAhead := grid[yAhead][xAhead]

		if stringAhead == "#" {
			direction = matrix.Multiply2d(turn90, direction)
			continue
		}

		if stringAhead == "." || stringAhead == "X" {
			x := position[0][0]
			y := position[1][0]

			grid[y][x] = "X"
		}

		position = matrix.Add2d(position, direction)
	}

	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == "X" {
				sum++
			}
		}
	}

	return strconv.Itoa(sum)
}
