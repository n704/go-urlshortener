package model

import (
	"net/url"

	valid "github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

// ShortLink is datatype
type ShortLink struct {
	gorm.Model
	URL     string `json:"url" valid:"url"`
	Shorten string `json:"shorten" gorm:"primary_key" valid:"optional"`
}

func (shortLink *ShortLink) Validate() error {
	_, err := valid.ValidateStruct(shortLink)
	if err != nil {
		return err
	}
	urlString := shortLink.URL
	u, err := url.ParseRequestURI(urlString)
	if err != nil {
		return err
	}
	shortLink.URL = u.String()
	return nil
}
