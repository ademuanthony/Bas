package main

import (
	"bitbucket.org/superfluxteam/pmsserver/common"
	"bitbucket.org/superfluxteam/pmsserver/routers"
	"github.com/codegangsta/negroni"
	"log"
	"net/http"
)

// Entry point of the program
func main() {
	//Calls startup logic
	common.StartUp()
	defer common.ShortDown()
	// Get mux router object
	router := routers.InitRoutes()
	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Printf("Listening at %s ... \n", common.AppConfig.Server)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	log.Println("stoping ...")
}
