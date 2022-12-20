package movie

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/movies").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/{id}", ctrl.GetById).Methods("GET")
	route.HandleFunc("", ctrl.Add).Methods("POST")
	route.HandleFunc("/{id}", ctrl.Update).Methods("PATCH")
	route.HandleFunc("/{id}", ctrl.Delete).Methods("DELETE")

}
