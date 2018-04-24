package model

import (
	"github.com/jinzhu/gorm"
)

// ShortLink is datatype
type ShortLink struct {
	gorm.Model
	URL     string `json:"url"`
	Shorten string `json:"shorten" gorm:"primary_key"`
}

//ShortLinks list of ShortLink
type ShortLinks []ShortLink
