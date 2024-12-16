package services

import (
	"advent-of-code-2024/pkg/file"
	"advent-of-code-2024/pkg/list"
	"strconv"
	"strings"
)

type IDay5Part2 interface {
	Solve() string
}

type Day5Part2 struct{}

func NewDay5Part2() IDay5Part2 {
	return &Day5Part2{}
}

func (obj *Day5Part2) Solve() string {
	// parse input
	fileString := file.ReadFileAsString()

	splitFileString := strings.Split(fileString, "\n")

	index := list.FirstIndex(splitFileString, "")

	rulesStrings := splitFileString[:index]
	updateStrings := splitFileString[index+1:]

	rules := make([][]int, len(rulesStrings))
	for index, rulesString := range rulesStrings {
		rulesStringSplit := strings.Split(rulesString, "|")

		if len(rulesStringSplit) != 2 {
			panic("error")
		}

		leftInt, err := strconv.Atoi(rulesStringSplit[0])
		if err != nil {
			panic(err)
		}
		rightInt, err := strconv.Atoi(rulesStringSplit[1])
		if err != nil {
			panic(err)
		}

		rules[index] = []int{leftInt, rightInt}
	}

	updates := make([][]int, len(updateStrings))
	for index, updateString := range updateStrings {
		updateStringSplit := strings.Split(updateString, ",")

		updates[index] = []int{}
		for _, partString := range updateStringSplit {
			partInt, err := strconv.Atoi(partString)
			if err != nil {
				panic(err)
			}
			updates[index] = append(updates[index], partInt)
		}
	}

	// solve
	afterMap := map[int][]int{}

	for _, rule := range rules {
		_, keyExists := afterMap[rule[0]]
		if !keyExists {
			afterMap[rule[0]] = []int{}
		}

		afterMap[rule[0]] = append(afterMap[rule[0]], rule[1])
	}

	sum := 0
	for _, update := range updates {
		if !IsOrdered(afterMap, update) {
			fixedUpdate := update

			i := len(fixedUpdate) - 1
			for !IsOrdered(afterMap, update) {

				theRest := fixedUpdate[:i]

				wasSwapped := false
				for restIndex, restElement := range theRest {
					for _, afterElement := range afterMap[fixedUpdate[i]] {
						if restElement == afterElement {
							swapOne := fixedUpdate[i]
							swapTwo := restElement

							fixedUpdate[i] = swapTwo
							fixedUpdate[restIndex] = swapOne

							wasSwapped = true
							break
						}
					}

					if wasSwapped {
						break
					}
				}

				i--

				if i < 0 {
					i = len(fixedUpdate) - 1
				}
			}

			sum += update[(len(fixedUpdate)-1)/2]
		}
	}

	return strconv.Itoa(sum)
}

func IsOrdered(afterMap map[int][]int, update []int) bool {
	for index, element := range update {
		shouldBeAfter := afterMap[element]
		whatIsBefore := update[:index]

		for _, after := range shouldBeAfter {
			for _, before := range whatIsBefore {
				if after == before {
					return false
				}
			}
		}
	}

	return true
}
