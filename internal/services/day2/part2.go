package services

import (
	"advent-of-code-2024/pkg/file"
	"advent-of-code-2024/pkg/list"
	"strconv"
	"strings"
)

type IDay2Part2 interface {
	Solve() string
}

type Day2Part2 struct{}

func NewDay2Part2() IDay2Part2 {
	return &Day2Part2{}
}

func (obj *Day2Part2) Solve() string {

	// parse file into list
	fileString := file.ReadFileAsString()

	splitFileString := strings.Split(fileString, "\n")

	lines := make([][]int, len(splitFileString))
	for linesIndex, lineString := range splitFileString {
		lineListString := strings.Split(lineString, " ")
		lineListString = list.Remove("", lineListString)
		line := make([]int, len(lineListString))
		for lineIndex, lineElementString := range lineListString {
			lineElementInt, err := strconv.Atoi(lineElementString)
			if err != nil {
				panic(err)
			}
			line[lineIndex] = lineElementInt
		}
		lines[linesIndex] = line
	}

	// solve
	safeCount := 0
	for _, line := range lines {
		isSafe := false

		if (IsAscendingArray(line) || IsDescendingArray(line)) && DiffIsBetween(1, 3, line) {
			isSafe = true
		}

		if !isSafe {
			for badLevelIndex, _ := range line {
				listWithBadLevelRemoved := list.Pop(badLevelIndex, line)
				if (IsAscendingArray(listWithBadLevelRemoved) || IsDescendingArray(listWithBadLevelRemoved)) && DiffIsBetween(1, 3, listWithBadLevelRemoved) {
					isSafe = true
					break
				}
			}
		}

		if isSafe {
			safeCount++
		}
	}

	return strconv.Itoa(safeCount)
}
