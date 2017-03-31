package routers

import (
	"github.com/gorilla/mux"
	"fmt"

	"github.com/ademuanthony/Bas/controllers"
)

func SetAuthRoute(router *mux.Router) *mux.Router {
	//applicationRoute := mux.NewRouter()

	fmt.Println("Setting application route")
	router.HandleFunc("/auth/register", controllers.AuthRegister).Methods("POST")
	router.HandleFunc("/auth/login", controllers.AuthLogin).Methods("POST")

	/*router.PathPrefix("/auth/register").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(applicationRoute),
	))*/
	return router
}
