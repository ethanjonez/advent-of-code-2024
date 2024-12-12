package services

import (
	"advent-of-code-2024/pkg/file"
	"advent-of-code-2024/pkg/stringhelper"
	"strconv"
)

type IDay3Part2 interface {
	Solve() string
}

type Day3Part2 struct{}

func NewDay3Part2() IDay3Part2 {
	return &Day3Part2{}
}

func (obj *Day3Part2) Solve() string {
	// parse file
	fileString := file.ReadFileAsString()

	// solve
	mulPattern := "mul"
	sum := 0

	for _, location := range stringhelper.GetSubStringLocations(fileString, mulPattern) {
		// locations are naturally ordered ascending

		firstNumber, secondNumber, isMulMatch := isMulString(fileString, location)
		if !isMulMatch {
			continue
		}

		mostRecentDoLocation := stringhelper.GetFirstSubStringLocation(fileString[:location], "do()", true)
		mostRecentDontLocation := stringhelper.GetFirstSubStringLocation(fileString[:location], "don't()", true)

		// covers all cases, since -1 is no match
		if mostRecentDoLocation < mostRecentDontLocation {
			continue
		}

		sum += firstNumber * secondNumber
	}

	return strconv.Itoa(sum)
}

func isMulString(fileString string, location int) (int, int, bool) {
	openBraceIndex := location + len("mul")

	if fileString[openBraceIndex] != '(' {
		return 0, 0, false
	}

	commaIndex := openBraceIndex + stringhelper.GetFirstSubStringLocation(fileString[openBraceIndex:], ",", false)

	if commaIndex < 0 {
		return 0, 0, false
	}

	closeBraceIndex := openBraceIndex + stringhelper.GetFirstSubStringLocation(fileString[openBraceIndex:], ")", false)

	if closeBraceIndex < 0 {
		return 0, 0, false
	}

	firstNumberString := fileString[openBraceIndex+1 : commaIndex]

	firstNumber, err := strconv.Atoi(firstNumberString)

	if err != nil {
		return 0, 0, false
	}

	secondNumberString := fileString[commaIndex+1 : closeBraceIndex]

	secondNumber, err := strconv.Atoi(secondNumberString)

	if err != nil {
		return 0, 0, false
	}

	return firstNumber, secondNumber, true
}
