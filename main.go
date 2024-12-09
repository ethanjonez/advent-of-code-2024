package main

import (
	"advent-of-code-2024/internal/routes"
	services "advent-of-code-2024/internal/services/day1"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	// register "services" and routes
	container.Provide(services.NewDay1Part1)
	container.Provide(services.NewDay1Part2)
	container.Provide(func() *gin.Engine {
		router := gin.Default()
		return router
	})
	container.Provide(routes.NewDay1)

	// host router, in future would have a list of background services in here too
	if err := container.Invoke(func(router *gin.Engine, day1 *routes.Day1) {
		day1.RegisterDay1(router)
		router.Run("localhost:8080")
	}); err != nil {
		panic(err)
	}
}
