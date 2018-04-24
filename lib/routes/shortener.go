package routes

import (
  "github.com/n704/go-urlshortener/lib/view"
	"github.com/gorilla/mux"
)

// RouteShortener addlist of urls
func RouteShortener(myRouter *mux.Router) {
	shortenerRoute := myRouter.PathPrefix("/shortener").Subrouter()
	shortenerRoute.Path("/urls").HandlerFunc(view.ListURL).Methods("GET")
	shortenerRoute.Path("/addUrl").HandlerFunc(view.AddURL).Methods("POST")
	shortenerRoute.Path("/deleteUrl/{shorten}").HandlerFunc(view.DeleteURL).Methods("GET")
	myRouter.HandleFunc("/{shorten}", view.RedirectURL).Methods("GET")
}
