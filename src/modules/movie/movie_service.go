package movie

import (
	"github.com/hendralatumeten/movieapp/src/database/models"
	"github.com/hendralatumeten/movieapp/src/interfaces"
	"github.com/hendralatumeten/movieapp/src/libs"
)

type movie_service struct {
	repo interfaces.MoveRepo
}

func NewService(reps interfaces.MoveRepo) *movie_service {
	return &movie_service{reps}
}

func (s *movie_service) GetAllMovie() *libs.Response {
	data, err := s.repo.FindAll()
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)
}

func (s *movie_service) GetById(id int) *libs.Response {
	data, err := s.repo.FindById(id)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)
}

func (s *movie_service) AddMovie(data *models.Movies) *libs.Response {
	data, err := s.repo.Save(data)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)
}
func (s *movie_service) UpdateMovie(id int, data *models.Movies) *libs.Response {
	data, err := s.repo.Update(id, data)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)
}

func (s *movie_service) DeleteMovie(id int) *libs.Response {
	data, err := s.repo.Delete(id)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)
}
