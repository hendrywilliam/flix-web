package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"flix/domain"
)

type moviesRepository struct{}

func NewMoviesCoreRepository() domain.MoviesRepository {
	return &moviesRepository{}
}

func (mr *moviesRepository) GetAllMovies() (movies []domain.Movie, err error) {
	apiKey := os.Getenv("TMDB_API_KEY")
	var tmdbResponse domain.TmdbMovieResponse
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://api.themoviedb.org/3/movie/now_playing?language=en-US&page=1", nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	res, _ := client.Do(req)
	err = json.NewDecoder(res.Body).Decode(&tmdbResponse)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return tmdbResponse.Results, nil
}
