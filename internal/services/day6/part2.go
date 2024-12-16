package services

type IDay6Part2 interface {
	Solve() string
}

type Day6Part2 struct{}

func NewDay6Part2() IDay6Part2 {
	return &Day6Part2{}
}

func (obj *Day6Part2) Solve() string {
	return ""
}
