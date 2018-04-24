package model

import (
	"github.com/jinzhu/gorm"
)

// ShortLink is datatype
type ShortLink struct {
	gorm.Model
	URL     string `json:"url" valid:"url"`
	Shorten string `json:"shorten" gorm:"primary_key" valid:"optional"`
}