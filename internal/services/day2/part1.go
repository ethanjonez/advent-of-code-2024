package services

import (
	"advent-of-code-2024/pkg/file"
	"advent-of-code-2024/pkg/list"
	"strconv"
	"strings"
)

type Day2Part1 struct{}

func NewDay2Part1() *Day2Part1 {
	return &Day2Part1{}
}

func (obj *Day2Part1) Solve() string {

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
		if (IsAscendingArray(line) || IsDescendingArray(line)) && DiffIsBetween(1, 3, line) {
			safeCount += 1
		}
	}

	return strconv.Itoa(safeCount)
}

func IsAscendingArray(listy []int) bool {
	for i := 0; i < len(listy)-1; i++ {
		if listy[i] >= listy[i+1] {
			return false
		}
	}

	return true
}

func IsDescendingArray(listy []int) bool {
	for i := 0; i < len(listy)-1; i++ {
		if listy[i] <= listy[i+1] {
			return false
		}
	}

	return true
}

func DiffIsBetween(min int, max int, listy []int) bool {
	for i := 0; i < len(listy)-1; i++ {
		diff := listy[i] - listy[i+1]

		if diff > 0 {
			if !(diff >= min && diff <= max) {
				return false
			}
		} else {
			if !(diff <= -min && diff >= -max) {
				return false
			}
		}
	}

	return true
}
