package services

import (
	"advent-of-code-2024/pkg/file"
	"advent-of-code-2024/pkg/list"
	"strconv"
	"strings"
)

type IDay5Part1 interface {
	Solve() string
}

type Day5Part1 struct{}

func NewDay5Part1() IDay5Part1 {
	return &Day5Part1{}
}

func (obj *Day5Part1) Solve() string {
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
		isOrdered := true
		for index, element := range update {
			shouldBeAfter := afterMap[element]
			whatIsBefore := update[:index]

			for _, after := range shouldBeAfter {
				for _, before := range whatIsBefore {
					if after == before {
						isOrdered = false
						break
					}
				}

				if !isOrdered {
					break
				}
			}

			if !isOrdered {
				break
			}
		}

		if isOrdered {
			sum += update[(len(update)-1)/2]
		}
	}

	return strconv.Itoa(sum)
}
