package matrix

func Multiply2d(a [][]int, b [][]int) [][]int { // need to extend for different slices - ideally probably want some separate type of struct to represent matrices
	if len(a[0]) != len(b) {
		panic("can't multiply matrices wrong dimensions")
	}

	// intialise c
	c := make([][]int, len(a))
	for index, _ := range c {
		c[index] = make([]int, len(b[0]))
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			rowA := a[i]
			colB := make([]int, len(b))
			for ind, rowB := range b {
				colB[ind] = rowB[j]
			}

			sum := 0
			for it := 0; it < len(colB); it++ {
				sum += rowA[it] * colB[it]
			}
			c[i][j] = sum
		}
	}

	return c
}

func Add2d(a [][]int, b [][]int) [][]int {
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		panic("can't add matrices wrong dimensions")
	}

	// initialise c
	c := make([][]int, len(a))
	for index, _ := range c {
		c[index] = make([]int, len(a[0]))
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}

	return c
}
