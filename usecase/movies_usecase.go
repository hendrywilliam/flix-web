package usecase

import "flix/domain"

type moviesUsecase struct {
	moviesRepository domain.MoviesRepository
}

func NewMoviesUseCase(moviesRepository domain.MoviesRepository) domain.MoviesUsecase {
	return &moviesUsecase{moviesRepository: moviesRepository}
}

func (mu *moviesUsecase) GetAllMovies() (movies []domain.Movie, err error) {
	return mu.moviesRepository.GetAllMovies()
}
