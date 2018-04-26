package model

import (
	valid "github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

// ShortLink is datatype
type ShortLink struct {
	gorm.Model
	URL     string `json:"url" valid:"url"`
	Shorten string `json:"shorten" gorm:"primary_key" valid:"optional"`
}

// Validate validate the ShortLink model and return error
func (shortLink *ShortLink) Validate() error {
	_, err := valid.ValidateStruct(shortLink)
	if err != nil {
		return err
	}
	return nil
}
