package services

import (
	"advent-of-code-2024/pkg/file"
	"strconv"
	"strings"
)

type IDay4Part2 interface {
	Solve() string
}

type Day4Part2 struct{}

func NewDay4Part2() IDay4Part2 {
	return &Day4Part2{}
}

func (obj *Day4Part2) Solve() string {
	// parse file
	fileString := file.ReadFileAsString()

	splitFileString := strings.Split(fileString, "\n")

	letterMatrix := make([][]string, len(splitFileString))
	for index, element := range splitFileString {
		letterMatrix[index] = strings.Split(element, "")
	}

	// solve
	sum := 0

	word := []string{
		"M",
		"S",
	}

	for i := 0; i < len(letterMatrix); i++ {
		for j := 0; j < len(letterMatrix[0]); j++ {
			if letterMatrix[i][j] != "A" {
				continue
			}
			// search
			directions := [][][]int{
				{{1, 1}, {-1, -1}},
				{{-1, -1}, {1, 1}},
				{{-1, 1}, {1, -1}},
				{{1, -1}, {-1, 1}},
			}

			matchCount := 0
			for _, directionCombo := range directions {
				isMatch := true

				for index, direction := range directionCombo {
					iToCheck := i + direction[0]
					jToCheck := j + direction[1]

					if iToCheck >= len(letterMatrix) ||
						jToCheck >= len(letterMatrix[0]) ||
						iToCheck < 0 ||
						jToCheck < 0 ||
						letterMatrix[iToCheck][jToCheck] != word[index] {
						isMatch = false
						break
					}
				}

				if isMatch {
					matchCount++
				}
			}

			if matchCount == 2 {
				sum++
			}
		}
	}

	return strconv.Itoa(sum)
}
