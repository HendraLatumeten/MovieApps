package movie

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hendralatumeten/movieapp/src/database/models"
	"github.com/hendralatumeten/movieapp/src/interfaces"
	"github.com/hendralatumeten/movieapp/src/libs"
)

type movie_ctrl struct {
	svc interfaces.MovieService
}

func NewCtrl(reps interfaces.MovieService) *movie_ctrl {
	return &movie_ctrl{svc: reps}
}

func (c *movie_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	c.svc.GetAllMovie().Send(w)
}

func (c *movie_ctrl) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(params)
	c.svc.GetById(id).Send(w)
}

func (c *movie_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var data models.Movies
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		libs.Respone(err.Error(), 401, true)
	}
	c.svc.AddMovie(&data).Send(w)
}

func (c *movie_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	var data models.Movies

	params := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(params)
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		libs.Respone(err.Error(), 401, true)
	}
	c.svc.UpdateMovie(id, &data).Send(w)
}

func (c *movie_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		libs.Respone(err.Error(), 401, true)
	}
	c.svc.DeleteMovie(id).Send(w)
}
