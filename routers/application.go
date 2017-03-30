package routers

import (
	"github.com/gorilla/mux"
	"github.com/ademuanthony/Bas/controllers"
	"github.com/codegangsta/negroni"
	"github.com/ademuanthony/Bas/common"
	"fmt"
)

func SetApplicationRoute(router *mux.Router) *mux.Router {
	applicationRoute := mux.NewRouter()

	fmt.Println("Setting application route")
	applicationRoute.HandleFunc("/applications", controllers.Create).Methods("POST")
	applicationRoute.HandleFunc("/applications", controllers.Create).Methods("GET")
	applicationRoute.HandleFunc("/applications/{id}", controllers.Create).Methods("GET")
	applicationRoute.HandleFunc("/applications/{id}", controllers.Create).Methods("DELETE")

	router.PathPrefix("/applications").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(applicationRoute),
	))
	return router
}
