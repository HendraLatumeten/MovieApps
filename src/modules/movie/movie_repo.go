package movie

import (
	"errors"

	"github.com/hendralatumeten/movieapp/src/database/models"
	"gorm.io/gorm"
)

type movie_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *movie_repo {
	return &movie_repo{db: db}
}

func (r *movie_repo) FindAll() (*models.MoviesAll, error) {
	var data models.MoviesAll

	result := r.db.Find(&data)

	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}

func (r *movie_repo) FindById(id int) (*models.Movies, error) {
	var data models.Movies

	result := r.db.First(&data, "id = ?", id)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}

func (r *movie_repo) Save(data *models.Movies) (*models.Movies, error) {
	result := r.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("gagal Menyimpan data")
	}

	return data, nil
}
func (r *movie_repo) Update(id int, data *models.Movies) (*models.Movies, error) {
	Idcheck := r.IdExsist(id)
	if !Idcheck {
		return nil, errors.New("id not found!")
	}
	result := r.db.Model(&data).Where("id = ?", id).Updates(&data)
	if result.Error != nil {
		return nil, errors.New("gagal Mengupdate data")
	}

	return data, nil
}

func (r *movie_repo) Delete(id int) (*models.Movies, error) {
	var data models.Movies
	Idcheck := r.IdExsist(id)
	if !Idcheck {
		return nil, errors.New("id not found!")
	}

	result := r.db.Where("id", id).Delete(&data)
	if result.Error != nil {
		return nil, errors.New("gagal Menghapus data")
	}
	result.Row()

	return &data, nil
}

func (r *movie_repo) IdExsist(id int) bool {
	var data models.Movies
	result := r.db.First(&data, "id=?", id)
	if result.Error != nil {
		return false
	}
	return true
}
