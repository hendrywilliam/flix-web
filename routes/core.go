package routes

import (
	"github.com/gin-gonic/gin"
)

func RoutesCore() *gin.Engine {
	r := gin.Default()

	NewMovieRoute(r)
	return r
}
