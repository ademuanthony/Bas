package routers

import (
	"github.com/ademuanthony/Bas/common"
	"github.com/ademuanthony/Bas/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetAuthRoute(router *mux.Router) *mux.Router {
	router.HandleFunc("/auth/login", controllers.AuthLogin).Methods("POST")

	applicationRoute := mux.NewRouter()
	applicationRoute.HandleFunc("/auth/register", controllers.AuthRegister).Methods("POST")

	router.PathPrefix("/auth/register").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(applicationRoute),
	))
	return router
}
