package services

type IDay5Part2 interface {
	Solve() string
}

type Day5Part2 struct{}

func NewDay5Part2() IDay5Part2 {
	return &Day5Part2{}
}

func (obj *Day5Part2) Solve() string {
	return ""
}
