/*
	Justrunit-Fury Service to manage docker containers in Suit instances
*/

package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/justrunit/docker"
)

func main() {

	/* Create router */
	router := mux.NewRouter()

	/* Add routes */
	router.HandleFunc("/run", docker.RunSnippet).Methods("POST")

	log.Println("Fury server listening on localhost:8081")

	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Println(err)
	}
}

