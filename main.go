package main

import (
	"advent-of-code-2024/internal/routes"
	day1 "advent-of-code-2024/internal/services/day1"
	day2 "advent-of-code-2024/internal/services/day2"
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

	// add routes
	handleErr(container.Provide(routes.NewDay1))
	handleErr(container.Provide(routes.NewDay2))
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
