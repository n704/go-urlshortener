package view

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
  "log"
	"time"
"github.com/n704/go-urlshortener/model"
 valid "github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//ListURL lising url
func ListURL(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	var shortLink []model.ShortLink
	db.Find(&shortLink)
	if err != nil {
		fmt.Printf(err.Error())
		panic("Error present")
	}
	json.NewEncoder(w).Encode(shortLink)
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

//AddURL adds new url to the lists
func AddURL(w http.ResponseWriter, r *http.Request) {
	// newLink := model.ShortLink{URL: "asd", Shorten: codeGenerator(8)}
	decoder := json.NewDecoder(r.Body)
	var newLink model.ShortLink
	err := decoder.Decode(&newLink)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	_, err = valid.ValidateStruct(&newLink)
  if err != nil {
    log.Printf(err.Error())
    json.NewEncoder(w).Encode("{\"data\":\""+err.Error()+"\"}")
    return
}
	if newLink.Shorten == "" {
		newLink.Shorten = codeGenerator(8)
	}
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	db.Create(&newLink)
	json.NewEncoder(w).Encode(newLink)
}

// DeleteURL delete shorten code from the links
func DeleteURL(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	link := model.ShortLink{}
	db.Where("shorten = ?", vars["shorten"]).First(&link)
	db.Delete(&link)
	json.NewEncoder(w).Encode("{\"data\":\"deleted successfully\"}")
}

//RedirectURL redirects to url if valid shorten code given
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	link := model.ShortLink{}
	db.Where("shorten = ?", vars["shorten"]).First(&link)
	url := link.URL
	if url != "" {
		http.Redirect(w, r, url, 301)
	} else {
		json.NewEncoder(w).Encode("{\"data\":\"invalid shorten code\"}")
	}
}
