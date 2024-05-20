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
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}
