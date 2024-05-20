package routes

import (
	"flix/controllers"
	"flix/repository"

	"github.com/julienschmidt/httprouter"
)

func NewMovieRoute(router *httprouter.Router) {
	mr := repository.NewMoviesCoreRepository()
	mc := &controllers.MoviesController{
		MoviesUsecase: mr,
	}

	router.GET("/movies", mc.GetMovies)
}
