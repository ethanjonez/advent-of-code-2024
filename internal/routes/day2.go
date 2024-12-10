package routes

import (
	services "advent-of-code-2024/internal/services/day2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Day2 struct {
	Day2Part1 services.IDay2Part1
	Day2Part2 services.IDay2Part2
}

func NewDay2(part1 services.IDay2Part1, part2 services.IDay2Part2) *Day2 {
	day2 := Day2{}
	day2.Day2Part1 = part1
	day2.Day2Part2 = part2

	return &day2
}

func (obj *Day2) RegisterDay2(route *gin.Engine) {
	day2group := route.Group("/2")
	{
		day2group.GET("/1", obj.part1Handler)
		day2group.GET("/2", obj.part2Handler)
	}
}

func (obj *Day2) part1Handler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, obj.Day2Part1.Solve())
}

func (obj *Day2) part2Handler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, obj.Day2Part2.Solve())
}
