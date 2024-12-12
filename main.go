package main

import (
	"advent-of-code-2024/internal/routes"
	day1 "advent-of-code-2024/internal/services/day1"
	day2 "advent-of-code-2024/internal/services/day2"
	day3 "advent-of-code-2024/internal/services/day3"
	day4 "advent-of-code-2024/internal/services/day4"
	day5 "advent-of-code-2024/internal/services/day5"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	// register "services" and routes
	handleErr(container.Provide(day1.NewDay1Part1))
	handleErr(container.Provide(day1.NewDay1Part2))
	handleErr(container.Provide(day2.NewDay2Part1))
	handleErr(container.Provide(day2.NewDay2Part2))
	handleErr(container.Provide(day3.NewDay3Part1))
	handleErr(container.Provide(day3.NewDay3Part2))
	handleErr(container.Provide(day4.NewDay4Part1))
	handleErr(container.Provide(day4.NewDay4Part2))
	handleErr(container.Provide(day5.NewDay5Part1))
	handleErr(container.Provide(day5.NewDay5Part2))

	// add routes
	handleErr(container.Provide(routes.NewDay1))
	handleErr(container.Provide(routes.NewDay2))
	handleErr(container.Provide(routes.NewDay3))
	handleErr(container.Provide(routes.NewDay4))
	handleErr(container.Provide(routes.NewDay5))
	handleErr(container.Provide(routes.NewServer))

	// host router, in future would have a list of background services in here too
	handleErr(container.Invoke(func(router *gin.Engine) {
		handleErr(router.Run("localhost:8080")) // auto injects routes in constructors
	}))
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
