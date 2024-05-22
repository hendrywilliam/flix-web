package routes

import (
	"flix/controllers"
	"flix/repository"
	"flix/usecase"

	"github.com/gin-gonic/gin"
)

func NewMovieRoute(r *gin.Engine) {
	mr := repository.NewMoviesCoreRepository()
	mc := &controllers.MoviesController{
		MoviesUsecase: usecase.NewMoviesUseCase(mr),
	}

	r.GET("/movies", mc.GetMovies)
	r.GET("/movie-genres", mc.GetMovieGenres)
}
