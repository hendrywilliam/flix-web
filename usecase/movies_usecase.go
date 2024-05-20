package usecase

import "flix/domain"

type moviesUsecase struct {
	moviesRepository domain.MoviesRepository
}

func NewMoviesUseCase(moviesRepository domain.MoviesRepository) *moviesUsecase {
	return &moviesUsecase{moviesRepository: moviesRepository}
}
