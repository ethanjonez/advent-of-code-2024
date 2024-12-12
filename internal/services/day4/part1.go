package services

import (
	"advent-of-code-2024/pkg/file"
	"strconv"
	"strings"
)

type IDay4Part1 interface {
	Solve() string
}

type Day4Part1 struct{}

func NewDay4Part1() IDay4Part1 {
	return &Day4Part1{}
}

func (obj *Day4Part1) Solve() string {
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
		"X",
		"M",
		"A",
		"S",
	}

	for i := 0; i < len(letterMatrix); i++ {
		for j := 0; j < len(letterMatrix[0]); j++ {
			// search
			directions := [][]int{
				{1, 1},
				{1, -1},
				{-1, 0},
				{0, -1},
				{-1, 1},
				{-1, -1},
				{0, 1},
				{1, 0},
			}

			for _, direction := range directions {
				isMatch := true

				for index, letter := range word {
					iToCheck := i + direction[0]*index
					jToCheck := j + direction[1]*index

					if iToCheck >= len(letterMatrix) ||
						jToCheck >= len(letterMatrix[0]) ||
						iToCheck < 0 ||
						jToCheck < 0 ||
						letterMatrix[iToCheck][jToCheck] != letter {
						isMatch = false
						break
					}
				}

				if isMatch {
					sum++
				}
			}
		}
	}

	return strconv.Itoa(sum)
}
