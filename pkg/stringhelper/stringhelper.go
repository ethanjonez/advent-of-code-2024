package stringhelper

func GetSubStringLocations(string string, match string) []int {
	var locations []int
	runeString := []rune(string) // why rune pls

	for index, _ := range runeString {
		for matchIndex, matchChar := range match {

			indexToCheck := index + matchIndex

			if indexToCheck >= len(string) {
				break
			}

			if runeString[indexToCheck] != matchChar {
				break
			}

			if matchIndex == len(match)-1 {
				locations = append(locations, index)
			}
		}
	}

	return locations
}

func GetFirstSubStringLocation(string string, match string, reverse bool) int {
	runeString := []rune(string) // why rune pls

	if reverse {
		for index := len(runeString) - 1; index >= 0; index-- {
			for matchIndex, matchChar := range match {

				indexToCheck := index + matchIndex

				if indexToCheck >= len(string) {
					break
				}

				if runeString[indexToCheck] != matchChar {
					break
				}

				if matchIndex == len(match)-1 {
					return index
				}
			}
		}
	} else {
		for index := 0; index < len(runeString); index++ {
			for matchIndex, matchChar := range match {

				indexToCheck := index + matchIndex

				if indexToCheck >= len(string) {
					break
				}

				if runeString[indexToCheck] != matchChar {
					break
				}

				if matchIndex == len(match)-1 {
					return index
				}
			}
		}
	}

	return -1
}
