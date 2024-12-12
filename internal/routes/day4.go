package routes

import (
	services "advent-of-code-2024/internal/services/day4"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Day4 struct {
	Day4Part1 services.IDay4Part1
	Day4Part2 services.IDay4Part2
}

func NewDay4(part1 services.IDay4Part1, part2 services.IDay4Part2) *Day4 {
	day4 := Day4{}
	day4.Day4Part1 = part1
	day4.Day4Part2 = part2

	return &day4
}

func (obj *Day4) RegisterDay4(route *gin.Engine) {
	day4group := route.Group("/4")
	{
		day4group.GET("/1", obj.part1Handler)
		day4group.GET("/2", obj.part2Handler)
	}
}

func (obj *Day4) part1Handler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, obj.Day4Part1.Solve())
}

func (obj *Day4) part2Handler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, obj.Day4Part2.Solve())
}
