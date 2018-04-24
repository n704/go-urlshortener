package test

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/n704/go-urlshortener/model"
	"github.com/n704/go-urlshortener/view"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
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

})
