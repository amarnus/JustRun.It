/*
	Justrunit-Potts Service to serve Stark Front-End
*/

//TODO: Write unit-tests for mux request handling with gocheck
//TODO: Fix Status Codes and Error Messsages according to HTTP Conventions

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/justrunit/models"
	"github.com/justrunit/routeinit"
	"log"
	"net/http"
)

type SnippetCreationResponse struct {
	Id string `json:"snippet_id"`
}

type SnippetsArrayResponse struct {
	Snippets *[]models.Snippet `json:"snippets"`
}

func main() {

	/* Create router */
	router := mux.NewRouter()

	/* Add routes */
	//GET /snippets?tag=<tag> - All snippets optionally filtered by tag.
	router.HandleFunc("/snippets", FilterSnippetsByTag).
		Methods("GET")
	//POST /snippets - Add a new snippet.
	router.HandleFunc("/snippets", CreateNewSnippet).
		Methods("POST")
	//GET /router/me - Snippets created by the requesting anonymous user.
	router.HandleFunc("/snippets/me", FilterSnippetsByUser).
		Methods("GET")

	//GET /snippet/<snippet_id> - Snippet code + metadata.
	router.HandleFunc("/snippet/{snippet_id}", FilterSnippetById).
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
		Methods("POST")
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
	validated, _, enc, _ := routeinit.InitHandling(req, resp, []string{})
	tag := req.FormValue("tag")
	if !validated {
		return
	}
	snippets, err := models.FindSnippetsByTag(tag)
	snippetArray := SnippetsArrayResponse{&snippets}
	if err != nil {
		enc.Encode(routeinit.ApiResponse{ErrorMessage: err.Error(), Status: false})
	} else {
		enc.Encode(routeinit.ApiResponse{Result: &snippetArray, Status: true})
	}
	return
}

func CreateNewSnippet(resp http.ResponseWriter, req *http.Request) {
	validated, _, enc, body := routeinit.InitHandling(req, resp, []string{
		"language",
	})
	if !validated {
		return
	}
	language := body["language"].(string)
	snippet := models.Snippet{LanguageCode: language}
	snippetId, ok, err := models.CreateSnippet(&snippet)
	if err != nil {
		enc.Encode(routeinit.ApiResponse{ErrorMessage: err.Error(), Status: ok})
	} else {
		ResponseResult := SnippetCreationResponse{snippetId}
		enc.Encode(routeinit.ApiResponse{Result: &ResponseResult, Status: ok})
	}
	return
}

//TODO: Refactor with anonymous session ids
func FilterSnippetsByUser(resp http.ResponseWriter, req *http.Request) {
	validated, _, enc, _ := routeinit.InitHandling(req, resp, []string{})
	sessionUser := req.FormValue("session_id")
	if !validated {
		return
	}
	snippets, err := models.FindSnippetsByUser(sessionUser)
	snippetArray := SnippetsArrayResponse{&snippets}
	if err != nil {
		enc.Encode(routeinit.ApiResponse{ErrorMessage: err.Error(), Status: false})
	} else {
		enc.Encode(routeinit.ApiResponse{Result: &snippetArray, Status: true})
	}
	return
}

func FilterSnippetById(resp http.ResponseWriter, req *http.Request) {
	validated, urlParams, enc, _ := routeinit.InitHandling(req, resp, []string{})
	if !validated {
		return
	}
	snippetId, ok := urlParams["snippet_id"]
	if ok != true {
		enc.Encode(routeinit.ApiResponse{ErrorMessage: "Invalid URL", Status: false})
		return
	}
	snippet, err := models.FindSnippetById(snippetId)
	if err != nil {
		enc.Encode(routeinit.ApiResponse{ErrorMessage: err.Error(), Status: ok})
	} else {
		enc.Encode(routeinit.ApiResponse{Result: &snippet, Status: true})
	}
	return
}

func UpdateSnippetById(resp http.ResponseWriter, req *http.Request) {
	routeVariables := mux.Vars(req)
	snippetId := routeVariables["snippet_id"]
	decoder := json.NewDecoder(req.Body)
	enc := json.NewEncoder(resp)
	snippet := models.Snippet{}
	err := decoder.Decode(&snippet)
	if err != nil {
		enc.Encode(routeinit.ApiResponse{ErrorMessage: "Malformed JSON Request", Status: false})
		return
	}
	ok, err := models.UpdateSnippetById(snippetId, &snippet)
	if err != nil {
		enc.Encode(routeinit.ApiResponse{ErrorMessage: err.Error(), Status: ok})
	} else {
		enc.Encode(routeinit.ApiResponse{Status: ok})
	}

}

func DeleteSnippetById(resp http.ResponseWriter, req *http.Request) {
	validated, urlParams, enc, _ := routeinit.InitHandling(req, resp, []string{})
	if !validated {
		return
	}
	snippetId, ok := urlParams["snippet_id"]
	if ok != true {
		enc.Encode(routeinit.ApiResponse{ErrorMessage: "Invalid URL", Status: false})
		return
	}
	_, err := models.DeleteSnippetById(snippetId)
	if err != nil {
		enc.Encode(routeinit.ApiResponse{ErrorMessage: err.Error(), Status: false})
	} else {
		enc.Encode(routeinit.ApiResponse{Status: true})
	}
	return
}

func RunSnippetById(resp http.ResponseWriter, req *http.Request) {
	validated, urlParams, enc, _ := routeinit.InitHandling(req, resp, []string{})
	if !validated {
		return
	}
	snippetId, ok := urlParams["snippet_id"]
	fmt.Printf("Running Snippet %s", snippetId)
	if ok != true {
		enc.Encode(routeinit.ApiResponse{ErrorMessage: "Invalid URL", Status: false})
		return
	}
	return
}

func LintSnippetById(resp http.ResponseWriter, req *http.Request) {
	validated, urlParams, enc, _ := routeinit.InitHandling(req, resp, []string{})
	if !validated {
		return
	}
	snippetId, ok := urlParams["snippet_id"]
	fmt.Printf("Linting Snippet %s", snippetId)
	if ok != true {
		enc.Encode(routeinit.ApiResponse{ErrorMessage: "Invalid URL", Status: false})
		return
	}
	return
}

func InstallDepsById(resp http.ResponseWriter, req *http.Request) {
	validated, urlParams, enc, _ := routeinit.InitHandling(req, resp, []string{})
	if !validated {
		return
	}
	snippetId, ok := urlParams["snippet_id"]
	fmt.Printf("Installing Dependencies Snippet %s", snippetId)
	if ok != true {
		enc.Encode(routeinit.ApiResponse{ErrorMessage: "Invalid URL", Status: false})
		return
	}
	return
}
