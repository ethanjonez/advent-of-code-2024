package file

import (
	"os"
)

func ReadFileAsString() string {
	dat, err := os.ReadFile("./puzzle.txt")

	if err != nil {
		panic(err)
	}

	return string(dat)
}
