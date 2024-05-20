package controllers

import (
	"encoding/json"
	"flix/domain"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MoviesController struct {
	MoviesUsecase domain.MoviesUsecase
}

func (mc *MoviesController) GetMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	movies, err := mc.MoviesUsecase.GetAllMovies()

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.ErrorResponse{
			Message: "Internal Server Error",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(domain.SuccessResponse[[]domain.Movie]{
		Message: "Successfully fetched movies.",
		Code:    http.StatusOK,
		Data:    movies,
	})
}
