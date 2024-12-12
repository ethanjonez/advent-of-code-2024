package services

import (
	"advent-of-code-2024/pkg/file"
	"advent-of-code-2024/pkg/stringhelper"
	"strconv"
)

type IDay3Part1 interface {
	Solve() string
}

type Day3Part1 struct{}

func NewDay3Part1() IDay3Part1 {
	return &Day3Part1{}
}

func (obj *Day3Part1) Solve() string {
	// parse file
	fileString := file.ReadFileAsString()

	// solve
	mulPattern := "mul"
	sum := 0

	for _, location := range stringhelper.GetSubStringLocations(fileString, mulPattern) {
		openBraceIndex := location + len(mulPattern)

		if fileString[openBraceIndex] != '(' {
			continue
		}

		commaIndex := openBraceIndex + stringhelper.GetFirstSubStringLocation(fileString[openBraceIndex:], ",", false)

		if commaIndex < 0 {
			continue
		}

		closeBraceIndex := openBraceIndex + stringhelper.GetFirstSubStringLocation(fileString[openBraceIndex:], ")", false)

		if closeBraceIndex < 0 {
			continue
		}

		firstNumberString := fileString[openBraceIndex+1 : commaIndex]

		firstNumber, err := strconv.Atoi(firstNumberString)

		if err != nil {
			continue
		}

		secondNumberString := fileString[commaIndex+1 : closeBraceIndex]

		secondNumber, err := strconv.Atoi(secondNumberString)

		if err != nil {
			continue
		}

		sum += firstNumber * secondNumber
	}

	return strconv.Itoa(sum)
}
