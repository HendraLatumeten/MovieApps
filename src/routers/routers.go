package routers

import (
	"errors"

	"github.com/gorilla/mux"
	"github.com/hendralatumeten/movieapp/src/database"
	"github.com/hendralatumeten/movieapp/src/modules/movie"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()
	db, err := database.Getconnection()
	if err != nil {
		return nil, errors.New("gagal init database")
	}
	movie.New(mainRoute, db)

	return mainRoute, nil
}
