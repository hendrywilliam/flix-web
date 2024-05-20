package routes

import (
	"github.com/julienschmidt/httprouter"
)

func RoutesCore() *httprouter.Router {
	r := httprouter.New()

	NewMovieRoute(r)
	return r
}
