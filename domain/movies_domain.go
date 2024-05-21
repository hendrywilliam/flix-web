package domain

type Dates struct {
	Maximum string `json:"maximum"`
	Minimum string `json:"minimum"`
}

type Movie struct {
	Adult            bool     `json:"adult"`
	BackdropPath     string   `json:"backdrop_path"`
	GenresId         []string `json:"genres_id"`
	ID               int      `json:"id"`
	OriginalLanguage string   `json:"original_language"`
	OringinalTitle   string   `json:"original_title"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	ReleaseDate      string   `json:"release_date"`
	Title            string   `json:"title"`
	Video            bool     `json:"video"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
}

type TmdbMovieResponse struct {
	Dates         Dates   `json:"dates"`
	Page          int     `json:"page"`
	Results       []Movie `json:"results"`
	Total_Pages   int     `json:"total_pages"`
	Total_Results int     `json:"total_results"`
}

type Genre struct {
	ID int `json:"id"`
}

type TmdbGenreResponse struct {
}

type MoviesRepository interface {
	GetAllMovies() ([]Movie, error)
}

type MoviesUsecase interface {
	GetAllMovies() ([]Movie, error)
}
