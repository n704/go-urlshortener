package routes

import (
	"github.com/gorilla/mux"
)

// AddRoutes add all routes to the server
func AddRoutes(myRouter *mux.Router) {

	RouteShortener(myRouter)

}
