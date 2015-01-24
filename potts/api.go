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
	//GET /snippets?tag=<tag>&page=<page> - All snippets optionally filtered by tag.
	router.HandleFunc("/snippets", FilterSnippetsByTag).
		Methods("GET")
	//POST /snippets - Add a new snippet.
	router.HandleFunc("/snippets", CreateNewSnippet).
		Methods("POST")
	//GET /router/me - Snippets created by the requesting anonymous user.
	router.HandleFunc("/snippets/me", FilterSnippetsByUser).
		Methods("GET")

	//GET /snippet/<snippet_id> - Snippet code + metadata.
	router.HandleFunc("/snippet/{snippet_id}", FilterSnippetsById).
		Methods("GET")
	//PUT /snippet/<snippet_id> - Update snippet code + metadata.
	router.HandleFunc("/snippet/{snippet_id}", UpdateSnippetById).
		Methods("PUT")
	//DELETE /snippet/<snippet_id> - Remove a snippet
	router.HandleFunc("/snippet/{snippet_id}", DeleteSnippetById).
		Methods("DELETE")
	//POST /snippet/<snippet_id>/run - Run the snippet
	router.HandleFunc("/snippet/{snippet_id}/run", RunSnippetById).
		Methods("POST")
	//POST /snippet/<snippet_id>/lint - Lint the snippet
	router.HandleFunc("/snippet/{snippet_id}/lint", LintSnippetById).
		Methods("GET")
	//POST /snippet/<snippet_id>/install - Install snippet dependencies
	router.HandleFunc("/snippet/{snippet_id}/install", InstallDepsById).
		Methods("POST")

	//Serves the static folder
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", router)
	log.Println("Potts waiting for Stark at localhost:80")
	err := http.ListenAndServe(":80", router)
	if err != nil {
		log.Println(err)
	}
}

func FilterSnippetsByTag(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "I'm %v under %s", req.URL.Path, req.Method)
	return
}

func CreateNewSnippet(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "I'm %v under %s", req.URL.Path, req.Method)
	return
}

func FilterSnippetsByUser(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "I'm %v under %s", req.URL.Path, req.Method)
	return
}

func FilterSnippetsById(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "I'm %v under %s", req.URL.Path, req.Method)
	return
}

func UpdateSnippetById(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "I'm %v under %s", req.URL.Path, req.Method)
	return
}

func DeleteSnippetById(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "I'm %v under %s", req.URL.Path, req.Method)
	return
}

func RunSnippetById(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "I'm %v under %s", req.URL.Path, req.Method)
	return
}

func LintSnippetById(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "I'm %v under %s", req.URL.Path, req.Method)
	return
}

func InstallDepsById(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "I'm %v under %s", req.URL.Path, req.Method)
	return
}
