package services

import (
	"advent-of-code-2024/pkg/file"
	"advent-of-code-2024/pkg/list"
	"strconv"
	"strings"
)

type Day1Part2 struct {
}

func NewDay1Part2() *Day1Part2 {
	return &Day1Part2{}
}

func (obj *Day1Part2) Solve() string {
	// parse file into list
	fileString := file.ReadFileAsString()

	splitFileString := strings.Split(fileString, "\n")

	leftSideValues := make([]int, len(splitFileString))
	rightSideValues := make([]int, len(splitFileString))

	for index, element := range splitFileString {
		splitLineString := strings.Split(element, " ")
		splitLineString = list.Remove("", splitLineString)
		leftSideValue, err := strconv.Atoi(splitLineString[0])
		if err != nil {
			panic(err)
		}
		rightSideValue, err := strconv.Atoi(splitLineString[1])
		if err != nil {
			panic(err)
		}
		leftSideValues[index] = leftSideValue
		rightSideValues[index] = rightSideValue
	}

	// solve
	sum := 0
	for _, leftSideValue := range leftSideValues {
		multiplier := list.Count(leftSideValue, rightSideValues)
		sum += leftSideValue * multiplier
	}

	return strconv.Itoa(sum)
}
