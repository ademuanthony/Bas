package main

import (
	"github.com/ademuanthony/Bas/common"
	"github.com/ademuanthony/Bas/routers"
	"github.com/codegangsta/negroni"
	"net/http"
	"log"
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
		Addr: common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening ...")
	err := server.ListenAndServe()
	if err != nil{
		log.Fatalf("%s\n", err)
	}
	log.Println("stoping ...")
}