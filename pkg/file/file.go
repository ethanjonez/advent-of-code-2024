package file

import (
	"os"
)

func ReadFileAsString() string {
	dat, err := os.ReadFile("/Users/ejones/Documents/GitHub/advent-of-code-2024/puzzle.txt")

	if err != nil {
		panic(err)
	}

	return string(dat)
}
