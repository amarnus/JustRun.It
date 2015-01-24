/*
	Justrunit-Potts Service to serve Stark Front-End
*/

package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	/* Create router */
	router := mux.NewRouter()

	/* Add routes */
	//Serves the static folder
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", router)

	log.Println("Potts waiting for Stark at localhost:80")
	err := http.ListenAndServe(":80", router)
	if err != nil {
		log.Println(err)
	}
}
