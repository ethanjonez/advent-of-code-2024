package services

import (
	"advent-of-code-2024/pkg/file"
	"advent-of-code-2024/pkg/list"
	"strconv"
	"strings"
)

type IDay1Part1 interface {
	Solve() string
}

type Day1Part1 struct{}

func NewDay1Part1() IDay1Part1 {
	return &Day1Part1{}
}

func (obj *Day1Part1) Solve() string {
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
	for i := 0; i < len(splitFileString); i++ {
		leftSideMinIndex, leftSideMinValue := list.GetLowestValue(leftSideValues)
		rightSideMinIndex, rightSideMinValue := list.GetLowestValue(rightSideValues)

		leftSideValues = list.Pop(leftSideMinIndex, leftSideValues)
		rightSideValues = list.Pop(rightSideMinIndex, rightSideValues)

		diff := rightSideMinValue - leftSideMinValue

		if diff < 0 {
			sum -= diff
		} else {
			sum += diff
		}
	}

	return strconv.Itoa(sum) // returns input for now
}
