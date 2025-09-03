package models

import (
	"gorm.io/gorm"
)

type Site struct {
	gorm.Model
	Name         string
	Url          string
	HttpStatus   int
	ResponseTime int
}
