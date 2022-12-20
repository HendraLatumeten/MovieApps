package interfaces

import (
	"github.com/hendralatumeten/movieapp/src/database/models"
	"github.com/hendralatumeten/movieapp/src/libs"
)

type MoveRepo interface {
	FindAll() (*models.MoviesAll, error)
	FindById(id int) (*models.Movies, error)
	Save(data *models.Movies) (*models.Movies, error)
	Update(id int, data *models.Movies) (*models.Movies, error)
	Delete(id int) (*models.Movies, error)
}

type MovieService interface {
	GetAllMovie() *libs.Response
	GetById(id int) *libs.Response
	AddMovie(data *models.Movies) *libs.Response
	UpdateMovie(id int, data *models.Movies) *libs.Response
	DeleteMovie(id int) *libs.Response
}
