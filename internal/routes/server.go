package routes

import "github.com/gin-gonic/gin"

// will expand as controllers expand - not much you can do about this
func NewServer(day1 *Day1, day2 *Day2, day3 *Day3, day4 *Day4, day5 *Day5) *gin.Engine {
	// intialise router
	router := gin.Default()

	// add top level routes
	day1.RegisterDay1(router)
	day2.RegisterDay2(router)
	day3.RegisterDay3(router)
	day4.RegisterDay4(router)
	day5.RegisterDay5(router)

	return router
}
