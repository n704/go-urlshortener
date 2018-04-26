package model

import (
	"math/rand"
	"time"

	valid "github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

// ShortLink is datatype
type ShortLink struct {
	gorm.Model
	URL     string `json:"url" valid:"url"`
	Shorten string `json:"shorten" gorm:"primary_key" valid:"optional"`
}

func codeGenerator(size int) string {
	var letters = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4",
		"5", "6", "7", "8", "9"}
	var shorten string
	shorten = ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= size; i++ {
		shorten += letters[rand.Intn(len(letters))]
	}
	return shorten
}

// Validate validate the ShortLink model and return error
func (shortLink *ShortLink) Validate() error {
	_, err := valid.ValidateStruct(shortLink)
	if err != nil {
		return err
	}
	if shortLink.Shorten == "" {
		shortLink.Shorten = codeGenerator(8)
	}
	return nil
}
