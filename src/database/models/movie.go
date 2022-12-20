package models

import "time"

type Movies struct {
	Id          int       `gorm:"type:int;primary_key;not null" json:"id"`
	Title       string    `gorm:"type:varchar;" json:"title"`
	Description string    `gorm:"type:varchar;" json:"description"`
	Rating      float64   `gorm:"type:float8;" json:"rating"`
	Image       string    `gorm:"type:varchar;" json:"image"`
	Created_at  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type MoviesAll []Movies
