package test

import (
	"fmt"
	"strings"

	"github.com/gorilla/mux"
	"github.com/n704/go-urlshortener/model"
	"github.com/n704/go-urlshortener/view"

	"net/http"
	"net/http/httptest"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Shortener", func() {
	var r *mux.Router
	model.InitialMigration()
	BeforeEach(func() {
		r = mux.NewRouter()
	})
	Describe("get Urls", func() {
		Context("get all urls", func() {
			It("should give code 200", func() {
				r.HandleFunc("/shorteners/urls", view.ListURL)
				server := httptest.NewServer(r)
				defer server.Close()
				url := fmt.Sprintf("%s/shorteners/urls", server.URL)
				r, err := http.NewRequest("GET", url, nil)
				res, err := http.DefaultClient.Do(r)
				if err != nil {
					fmt.Printf(err.Error())
				}
				Expect(res.StatusCode).To(Equal(200))
			})
		})
	})

	Describe("Add Url", func() {
		Context("Add new url to the list", func() {
			It("should give code 201", func() {
				r.HandleFunc("/shorteners/addurl", view.AddURL)
				server := httptest.NewServer(r)
				defer server.Close()
				url := fmt.Sprintf("%s/shorteners/addurl", server.URL)
				userData := `{"url": "http://www.facebook.com"}`
				r, err := http.NewRequest("POST", url, strings.NewReader(userData))
				res, err := http.DefaultClient.Do(r)
				if err != nil {
					fmt.Printf(err.Error())
				}
				Expect(res.StatusCode).To(Equal(201))
			})
			It("should give code 400", func() {
				r.HandleFunc("/shorteners/addurl", view.AddURL)
				server := httptest.NewServer(r)
				defer server.Close()
				url := fmt.Sprintf("%s/shorteners/addurl", server.URL)
				userData := `{"url": "ceboo"}`
				r, err := http.NewRequest("POST", url, strings.NewReader(userData))
				res, err := http.DefaultClient.Do(r)
				if err != nil {
					Expect(err.Error()).To(Equal("ceboo does not validate as url"))
				}
				Expect(res.StatusCode).To(Equal(400))
			})
		})
	})

	Describe("Url Redirect", func() {
		Context("Redirect testing for shorten code", func() {
			It("give 200 status for valid shorten code.", func() {
				r.HandleFunc("/shorteners/addurl", view.AddURL)
				r.HandleFunc("/shorteners/{shorten}", view.RedirectURL)
				server := httptest.NewServer(r)
				defer server.Close()
				url := fmt.Sprintf("%s/shorteners/addurl", server.URL)
				userData := `{"url": "https://facebook.com", "shorten": "db.com1"}`
				r, err := http.NewRequest("POST", url, strings.NewReader(userData))
				_, err = http.DefaultClient.Do(r)
				if err != nil {
					fmt.Printf(err.Error())
				}
				url = fmt.Sprintf("%s/shorteners/db.com1", server.URL)
				r, err = http.NewRequest("GET", url, nil)
				res, err := http.DefaultClient.Do(r)
				if err != nil {
					fmt.Printf(err.Error())
				}
				Expect(res.Request.URL.String()).To(Equal("https://www.facebook.com/"))
				Expect(res.StatusCode).To(Equal(200))
			})
		})
	})
})
