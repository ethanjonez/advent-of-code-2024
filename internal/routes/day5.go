package routes

import (
	services "advent-of-code-2024/internal/services/day5"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Day5 struct {
	Day5Part1 services.IDay5Part1
	Day5Part2 services.IDay5Part2
}

func NewDay5(part1 services.IDay5Part1, part2 services.IDay5Part2) *Day5 {
	day5 := Day5{}
	day5.Day5Part1 = part1
	day5.Day5Part2 = part2

	return &day5
}

func (obj *Day5) RegisterDay5(route *gin.Engine) {
	day5group := route.Group("/5")
	{
		day5group.GET("/1", obj.part1Handler)
		day5group.GET("/2", obj.part2Handler)
	}
}

func (obj *Day5) part1Handler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, obj.Day5Part1.Solve())
}

func (obj *Day5) part2Handler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, obj.Day5Part2.Solve())
}
