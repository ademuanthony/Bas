package routers

import (
	"github.com/ademuanthony/bas/common"
	"github.com/ademuanthony/bas/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetAclRoutes(router *mux.Router) *mux.Router {
	applicationRoute := mux.NewRouter()

	applicationRoute.HandleFunc("/resources", controllers.CreateResource).Methods("POST")
	applicationRoute.HandleFunc("/resources/createmany", controllers.CreateResources).Methods("POST")
	applicationRoute.HandleFunc("/resources", controllers.GetResources).Methods("GET")
	applicationRoute.HandleFunc("/resources/{id}", controllers.GetResourceById).Methods("GET")
	applicationRoute.HandleFunc("/resources/{id}/delete", controllers.DeleteResource).Methods("POST")
	applicationRoute.HandleFunc("/resources/{id}/update", controllers.UpdateResource).Methods("POST")

	router.PathPrefix("/resources").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(applicationRoute),
	))

	router = setAclRoleRoute(router)
	router = setAclUserRoute(router)

	return router
}

func setAclRoleRoute(router *mux.Router) *mux.Router {
	applicationRoute := mux.NewRouter()

	applicationRoute.HandleFunc("/roles", controllers.CreateRole).Methods("POST")
	applicationRoute.HandleFunc("/roles", controllers.GetRoles).Methods("GET")
	applicationRoute.HandleFunc("/roles/{id}", controllers.GetRole).Methods("GET")
	applicationRoute.HandleFunc("/roles/{id}/delete", controllers.DeleteRole).Methods("POST")

	applicationRoute.HandleFunc("/roles/{roleId}/resources", controllers.GetResourceInRole).Methods("GET")
	applicationRoute.HandleFunc("/roles/{roleId}/resources/{resourceId}", controllers.AddResourceToRole).Methods("POST")
	applicationRoute.HandleFunc("/roles/{roleId}/resources", controllers.AddResourcesToRole).Methods("POST")
	applicationRoute.HandleFunc("/roles/{roleId}/resources/{resourceId}/delete", controllers.RemoveResourceFromRole).Methods("POST")

	applicationRoute.HandleFunc("/roles/{roleId}/users", controllers.GetUsersInRole).Methods("GET")
	applicationRoute.HandleFunc("/roles/{roleId}/users/{userId}", controllers.AddUserToRole).Methods("POST")
	applicationRoute.HandleFunc("/roles/{roleId}/users/{userId}/delete", controllers.RemoveUserFromRole).Methods("POST")

	router.PathPrefix("/roles").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(applicationRoute),
	))
	return router
}

func setAclUserRoute(router *mux.Router) *mux.Router {
	applicationRoute := mux.NewRouter()

	applicationRoute.HandleFunc("/users/{userId}/roles", controllers.GetRolesForUser).Methods("GET")
	applicationRoute.HandleFunc("/users/{userId}/resources", controllers.GetResourceForUser).Methods("GET")


	router.PathPrefix("/users/").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(applicationRoute),
	))
	return router
}
