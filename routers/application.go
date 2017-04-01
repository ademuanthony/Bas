package routers

import (
	"fmt"
	"github.com/ademuanthony/Bas/common"
	"github.com/ademuanthony/Bas/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetApplicationRoute(router *mux.Router) *mux.Router {
	applicationRoute := mux.NewRouter()

	fmt.Println("Setting application route")
	applicationRoute.HandleFunc("/applications", controllers.ApplicationCreate).Methods("POST")
	applicationRoute.HandleFunc("/applications", controllers.ApplicationCreate).Methods("GET")
	applicationRoute.HandleFunc("/applications/{id}", controllers.ApplicationCreate).Methods("GET")
	applicationRoute.HandleFunc("/applications/{id}", controllers.ApplicationCreate).Methods("DELETE")

	router.PathPrefix("/applications").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(applicationRoute),
	))
	return router
}
