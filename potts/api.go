/*
	Justrunit-Potts Service to serve Stark Front-End
*/

//TODO: Write unit-tests for mux request handling with gocheck
//TODO: Fix Status Codes and Error Messsages according to HTTP Conventions
package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/justrunit/models"
	"github.com/justrunit/routeinit"
	"gopkg.in/boj/redistore.v1"
	"log"
	"net/http"
	"time"
)

type SnippetCreationResponse struct {
	Id string `json:"snippet_id"`
}

type SnippetsArrayResponse struct {
	Snippets *[]models.Snippet `json:"snippets"`
}

var globalStore *redistore.RediStore

const MaxAge int = 4 * 30 * 24 * 3600

func main() {

	/* Create router */
	router := mux.NewRouter()
	store, err := redistore.NewRediStore(10, "tcp", "172.17.42.1:6379", "", []byte("secret-key"))
	if err != nil {
		panic(err)
	}
	store.SetMaxAge(MaxAge)
	globalStore = store
	defer store.Close()
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
	//POST /snippet/<snippet_id>/fork - Forks an existing snippe.t
	router.HandleFunc("/snippet/{snippet_id}/fork", ForkSnippetById).
		Methods("POST")

	//Serves the static folder
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", router)
	log.Println("Potts waiting for Stark at localhost:8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Println(err)
	}
}

func setACLHeaders(resp http.ResponseWriter) http.ResponseWriter {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
	return resp
}

func setSessionIDAsCookie(resp http.ResponseWriter, sessionId string) http.ResponseWriter {
	d := time.Duration(MaxAge) * time.Second
	Expires := time.Now().Add(d)
	http.SetCookie(resp, &http.Cookie{Name: "session_id", Value: sessionId, MaxAge: MaxAge, Path: "/", Expires: Expires})
	return resp
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
		"language_code",
	})
	if !validated {
		return
	}
	session, _ := globalStore.Get(req, "session-name")
	session.Save(req, resp)
	resp = setSessionIDAsCookie(resp, session.ID)
	language := body["language_code"].(string)
	snippet := models.Snippet{LanguageCode: language, Tags: []string{language}, SessionId: session.ID}
	snippetId, ok, err := models.CreateSnippet(&snippet)
	if err != nil {
		enc.Encode(routeinit.ApiResponse{ErrorMessage: err.Error(), Status: ok})
	} else {
		ResponseResult := SnippetCreationResponse{snippetId}
		enc.Encode(routeinit.ApiResponse{Result: &ResponseResult, Status: ok})
	}
	return
}

func FilterSnippetsByUser(resp http.ResponseWriter, req *http.Request) {
	validated, _, enc, _ := routeinit.InitHandling(req, resp, []string{})
	session, _ := globalStore.Get(req, "session-name")
	sessionUser := session.ID
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
	session, _ := globalStore.Get(req, "session-name")
	session.Save(req, resp)
	resp = setSessionIDAsCookie(resp, session.ID)
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

func ForkSnippetById(resp http.ResponseWriter, req *http.Request) {
	routeVariables := mux.Vars(req)
	snippetId := routeVariables["snippet_id"]
	session, _ := globalStore.Get(req, "session-name")
	session.Save(req, resp)
	resp = setSessionIDAsCookie(resp, session.ID)
	snippet, err := models.FindSnippetById(snippetId)
	snippet.SessionId = session.ID
	newSnippetId, ok, err := models.CreateSnippet(snippet)
	enc := json.NewEncoder(resp)
	if err != nil {
		enc.Encode(routeinit.ApiResponse{ErrorMessage: err.Error(), Status: ok})
	} else {
		ResponseResult := SnippetCreationResponse{newSnippetId}
		enc.Encode(routeinit.ApiResponse{Result: &ResponseResult, Status: ok})
	}
	return
}
