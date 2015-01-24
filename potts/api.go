/*
	Justrunit-Potts Service to serve Stark Front-End
*/

package main

import (
	"fmt"
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

	log.Println("Potts waiting for Stark at localhost:5000")
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		log.Println(err)
	}
}
