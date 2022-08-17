package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title string
	Info  string
}
