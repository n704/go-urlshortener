package view

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/n704/go-urlshortener/model"
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
	linkPointer := &newLink
	err = linkPointer.Validate()
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("{\"data\":\"" + err.Error() + "\"}")
		return
	}
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	db.Create(&newLink)
	w.WriteHeader(201)
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
		// w.WriteHeader(302)
		http.Redirect(w, r, url, 301)
	} else {
		json.NewEncoder(w).Encode("{\"data\":\"invalid shorten code\"}")
	}
}
