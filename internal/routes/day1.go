package routes

import (
	services "advent-of-code-2024/internal/services/day1"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Day1 struct {
	Day1Part1 services.Day1Part1
}

type response struct {
	Answer         string  `json:"answer"`
	ElapsedSeconds float64 `json:"elapsed"`
}

func NewDay1(part1 *services.Day1Part1) *Day1 {
	day1 := Day1{}
	day1.Day1Part1 = *part1
	return &day1
}

func (obj *Day1) RegisterDay1(router *gin.Engine) {
	day1Group := router.Group("/1")
	{
		day1Group.GET("/1", obj.part1Handler)
	}
}

func (obj *Day1) part1Handler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, obj.Day1Part1.Solve("test"))
}