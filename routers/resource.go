package routers

import (
	"fmt"
	"github.com/ademuanthony/Bas/common"
	"github.com/ademuanthony/Bas/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetResourceRoute(router *mux.Router) *mux.Router {
	applicationRoute := mux.NewRouter()

	fmt.Println("Setting resource route")
	applicationRoute.HandleFunc("/resources", controllers.CreateResource).Methods("POST")
	applicationRoute.HandleFunc("/resources", controllers.GetResources).Methods("GET")
	applicationRoute.HandleFunc("/resources/{id}", controllers.GetResource).Methods("GET")
	applicationRoute.HandleFunc("/resources/{id}", controllers.DeleteResource).Methods("DELETE")

	router.PathPrefix("/resources").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(applicationRoute),
	))
	return router
}
