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
	container.Provide(day1.NewDay1Part1)
	container.Provide(day1.NewDay1Part2)
	container.Provide(day2.NewDay2Part1)
	container.Provide(day2.NewDay2Part2)
	container.Provide(func() *gin.Engine {
		router := gin.Default()
		return router
	})
	container.Provide(routes.NewDay1)
	container.Provide(routes.NewDay2)

	// host router, in future would have a list of background services in here too
	if err := container.Invoke(func(router *gin.Engine, day1 *routes.Day1, day2 *routes.Day2) {
		day1.RegisterDay1(router)
		day2.RegisterDay2(router)
		router.Run("localhost:8080")
	}); err != nil {
		panic(err)
	}
}
