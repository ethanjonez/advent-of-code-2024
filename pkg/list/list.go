package list

// index, element
func GetLowestValue(listy []int) (int, int) {
	if len(listy) < 1 {
		panic("list is empty")
	}

	minIndex := 0
	minElement := listy[minIndex]
	for index, element := range listy {
		if element < minElement {
			minElement = element
			minIndex = index
		}
	}

	return minIndex, minElement
}

func Pop[V int | string](index int, listy []V) (newListy []V) {
	if index >= len(listy) || index < 0 {
		panic("index out of range for list")
	}

	newListy = append(listy[:index], listy[index+1:]...)

	return newListy
}

func Remove[V int | string](element V, listy []V) []V {
	for {
		noChanges := true
		for index, listyElement := range listy {
			if listyElement == element {
				listy = Pop(index, listy)
				noChanges = false
			}
		}

		if noChanges {
			break
		}
	}

	return listy
}

func Count[V comparable](element V, listy []V) int {
	count := 0

	for _, listyElement := range listy {
		if element == listyElement {
			count++
		}
	}

	return count
}
