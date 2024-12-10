package routes

import "github.com/gin-gonic/gin"

// will expand as controllers expand - not much you can do about this
func NewServer(day1 *Day1, day2 *Day2) *gin.Engine {
	// intialise router
	router := gin.Default()

	// add top level routes
	day1.RegisterDay1(router)
	day2.RegisterDay2(router)

	return router
}
