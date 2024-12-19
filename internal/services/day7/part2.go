package day7

import (
	"advent-of-code-2024/pkg/file"
	"advent-of-code-2024/pkg/list"
	"strconv"
	"strings"
)

func solve() string {

	// parse data
	fileString := file.ReadFileAsString()

	splitFileString := strings.Split(fileString, "\n")

	data := map[int][]int{}

	for _, line := range splitFileString {
		lineSplit := strings.Split(line, ":")

		total, err := strconv.Atoi(lineSplit[0])

		if err != nil {
			panic(err)
		}

		rhsSplit := strings.Split(lineSplit[1], " ")

		rhsSplitClean := list.Remove("", rhsSplit)

		data[total] = make([]int, len(rhsSplitClean))

		for index, strElement := range rhsSplitClean {
			intElement, err := strconv.Atoi(strElement)

			if err != nil {
				panic(err)
			}

			data[total][index] = intElement
		}
	}

	// solves
	sum := 0
	for total, numbers := range data {
		if recurse(total, numbers, []string{}, []string{"+", "*", "||"}) {
			sum += total
		}
	}

	return strconv.Itoa(sum)
}

func recurse(total int, values []int, valueOperations []string, operations []string) bool {
	if len(valueOperations)+1 == len(values) {
		actualTotal := compute(values, valueOperations)

		if actualTotal == total {
			return true
		} else {
			return false
		}
	}

	for _, operation := range operations {
		newValueOfOperations := append(valueOperations, operation)
		isMatch := recurse(total, values, newValueOfOperations, operations)

		if isMatch == true {
			return true
		}
	}

	return false
}

func compute(values []int, valueOperations []string) int {
	total := values[0]
	for i, operation := range valueOperations {
		switch operation {
		case "*":
			total *= values[i+1]
		case "+":
			total += values[i+1]
		case "||":
			totalTemp, err := strconv.Atoi(strconv.Itoa(total) + strconv.Itoa(values[i+1]))

			if err != nil {
				panic(err)
			}

			total = totalTemp
		default:
			panic("not a proper operation")
		}
	}

	return total
}
