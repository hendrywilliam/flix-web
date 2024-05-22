package controllers

import (
	"flix/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MoviesController struct {
	MoviesUsecase domain.MoviesUsecase
}

func (mc *MoviesController) GetMovies(c *gin.Context) {
	movies, err := mc.MoviesUsecase.GetAllMovies()

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Internal Server Error",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse[[]domain.Movie]{
		Message: "Successfully fetched movies.",
		Code:    http.StatusOK,
		Data:    movies,
	})
}

func (mc *MoviesController) GetMovieGenres(c *gin.Context) {
	genres, err := mc.MoviesUsecase.GetMovieGenres()

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Internal Server Error",
			Code:    http.StatusInternalServerError,
		})
	}

	c.JSON(http.StatusOK, domain.SuccessResponse[[]domain.Genre]{
		Message: "Successfully fetched movie genres.",
		Code:    http.StatusOK,
		Data:    genres,
	})
}
