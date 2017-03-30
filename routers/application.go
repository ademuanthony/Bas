package routers

import (
	"github.com/gorilla/mux"
	"github.com/ademuanthony/Bas/controllers"
	"github.com/codegangsta/negroni"
	"github.com/ademuanthony/Bas/common"
)

func SetApplicationRoute(router *mux.Router) *mux.Router {
	applicationRoute := mux.NewRouter()

	applicationRoute.HandleFunc("/applications/create", controllers.Create).Methods("POST")

	router.PathPrefix("applications").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(applicationRoute),
	))
	return router
}
