/*
	Justrunit-Fury Service to manage docker containers in Suit instances
*/

package main

import (
	"net/http"
	"log"
	//"github.com/gorilla/mux"
	"github.com/justrunit/docker"
	"github.com/justrunit/furywebsockets"
)

func main() {

	/* Create router */
	//router := mux.NewRouter()

	/* Add routes */
	http.HandleFunc("/run/complete", docker.RunSnippetSync)
	http.HandleFunc("/run", docker.RunSnippetAsync)
	http.HandleFunc("/ws/io", furywebsockets.ServeWs);

	log.Println("Fury server listening on localhost:8081")

	/* HTTP Handler */
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Println(err)
	}

}

