package routes

import (
	services "advent-of-code-2024/internal/services/day3"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Day3 struct {
	Day3Part1 services.IDay3Part1
	Day3Part2 services.IDay3Part2
}

func NewDay3(part1 services.IDay3Part1, part2 services.IDay3Part2) *Day3 {
	day3 := Day3{}
	day3.Day3Part1 = part1
	day3.Day3Part2 = part2

	return &day3
}

func (obj *Day3) RegisterDay3(route *gin.Engine) {
	day3group := route.Group("/3")
	{
		day3group.GET("/1", obj.part1Handler)
		day3group.GET("/2", obj.part2Handler)
	}
}

func (obj *Day3) part1Handler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, obj.Day3Part1.Solve())
}

func (obj *Day3) part2Handler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, obj.Day3Part2.Solve())
}
