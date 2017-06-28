package routers

import (
	"bitbucket.org/superfluxteam/pmsserver/controllers"
	"github.com/gorilla/mux"
)

func SetAuthRoute(router *mux.Router) *mux.Router {

	router.HandleFunc("/auth/changepassword", controllers.ChangePassword).Methods("POST")

	return router
}
