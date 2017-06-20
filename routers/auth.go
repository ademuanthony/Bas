package routers

import (
	"github.com/ademuanthony/Bas/common"
	"github.com/ademuanthony/Bas/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetAuthRoute(router *mux.Router) *mux.Router {
	router.HandleFunc("/auth/login", controllers.AuthLogin).Methods("POST")
	router.HandleFunc("/utils/sendmail", controllers.SendEmail).Methods("POST")
	router.HandleFunc("/utils/sendmails", controllers.SendMultipleEmail).Methods("POST")

	applicationRoute := mux.NewRouter()

	applicationRoute.HandleFunc("/auth/register", controllers.AuthRegister).Methods("POST")
	applicationRoute.HandleFunc("/auth/{id}/update", controllers.AuthUpdate).Methods("POST")
	applicationRoute.HandleFunc("/auth/createaccounts", controllers.AuthCreateAccounts).Methods("POST")


	//applicationRoute.HandleFunc("/auth/register", controllers.AuthRegister).Methods("POST")
	applicationRoute.HandleFunc("/auth/changepassword", controllers.ChangePassword).Methods("POST")
	applicationRoute.HandleFunc("/auth/{id}/changepassword", controllers.ChangePasswordForId).Methods("POST")

	router.PathPrefix("/auth/changepassword").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(applicationRoute),
	))


	router.PathPrefix("/auth/register").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(applicationRoute),
	))
	router.PathPrefix("/auth/{id}/update").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(applicationRoute),
	))
	router.PathPrefix("/auth/createaccounts").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(applicationRoute),
	))

	router.PathPrefix("/auth/{id}/changepassword").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(applicationRoute),
	))
	return router
}
