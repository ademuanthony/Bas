package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	//Routes for application
	router = SetApplicationRoute(router)
	router = SetAuthRoute(router)

	return router
}
