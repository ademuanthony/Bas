package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	//Routes for application
	router = SetAuthRoute(router)

	return router
}
