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
		for index, _ := range update {
			if index == 0 {
				continue // assuming there isn't any one page updates
			}

			_ = ContainsElement(update[index-1], afterMap, update[index])

			if !ContainsElement(update[index], afterMap, update[index-1]) || ContainsElement(update[index-1], afterMap, update[index]) {
				isOrdered = false
				break
			}
		}

		if isOrdered {
			sum += update[(len(update)-1)/2]
		}
	}

	return strconv.Itoa(sum)
}

// returns is before
func ContainsElement(element int, afterMap map[int][]int, beforeElement int) bool {
	_, beforeElementExists := afterMap[beforeElement]
	if !beforeElementExists {
		return false
	}

	for _, afterElement := range afterMap[beforeElement] {
		if afterElement == element {
			return true
		}
	}

	anyTrue := false
	for _, afterElement := range afterMap[beforeElement] {
		if ContainsElement(element, afterMap, afterElement) {
			anyTrue = true
			break
		}
	}

	return anyTrue
}
