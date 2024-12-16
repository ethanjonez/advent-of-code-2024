package routes

import (
	services "advent-of-code-2024/internal/services/day6"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Day6 struct {
	Day6Part1 services.IDay6Part1
	Day6Part2 services.IDay6Part2
}

func NewDay6(part1 services.IDay6Part1, part2 services.IDay6Part2) *Day6 {
	day6 := Day6{}
	day6.Day6Part1 = part1
	day6.Day6Part2 = part2

	return &day6
}

func (obj *Day6) RegisterDay6(route *gin.Engine) {
	day6group := route.Group("/6")
	{
		day6group.GET("/1", obj.part1Handler)
		day6group.GET("/2", obj.part2Handler)
	}
}

func (obj *Day6) part1Handler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, obj.Day6Part1.Solve())
}

func (obj *Day6) part2Handler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, obj.Day6Part2.Solve())
}
