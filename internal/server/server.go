package server

import "github.com/gorilla/mux"

func initializeRoutes(r *mux.Router) {

}

func Setup() *mux.Router {
	r := mux.NewRouter()

	initializeRoutes(r)
	return r
}
