package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Snippet struct {
	LanguageCode string   `json:"language_code"`
	SessionId    string   `json:"session_id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Tags         []string `json:"tags"`
	Code         string   `json:"code"`
	Deps         []string `json:"deps"`
}

func getCollection(db string, collection string) (session *mgo.Session, c *mgo.Collection, err error) {
	session, err = mgo.Dial("172.17.42.1")
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c = session.DB(db).C(collection)
	return
}

func CreateSnippet(snippet *Snippet) (ok bool, err error) {
	session, c, _ := getCollection("user_snippets", "snippets")
	defer session.Close()
	err = c.Insert(snippet)
	if err != nil {
		log.Fatal(err)
		ok = false
		return
	}
	ok = true
	return
}

func FindSnippetById(snippetId string) (snippet *Snippet, err error) {
	snippet = &Snippet{}
	session, c, _ := getCollection("user_snippets", "snippets")
	defer session.Close()
	err = c.FindId(bson.ObjectIdHex(snippetId)).One(snippet)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func UpdateSnippetById(snippetId string, snippet *Snippet) (ok bool, err error) {
	session, c, _ := getCollection("user_snippets", "snippets")
	defer session.Close()
	err = c.UpdateId(bson.ObjectIdHex(snippetId), snippet)
	ok = true
	if err != nil {
		log.Fatal(err)
		ok = false
		return
	}
	return
}

func FindSnippetsByUser(snippetUser string) (snippets []Snippet, err error) {
	snippets = []Snippet{}
	session, c, _ := getCollection("user_snippets", "snippets")
	defer session.Close()
	err = c.Find(bson.M{"sessionid": snippetUser}).Iter().All(&snippets)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func FindSnippetsByTag(snippetTag string) (snippets []Snippet, err error) {
	snippets = []Snippet{}
	session, c, _ := getCollection("user_snippets", "snippets")
	defer session.Close()
	err = c.Find(bson.M{"tags": snippetTag}).Iter().All(&snippets)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func DeleteSnippetById(snippetId string) (ok bool, err error) {
	session, c, _ := getCollection("user_snippets", "snippets")
	defer session.Close()
	err = c.RemoveId(bson.ObjectIdHex(snippetId))
	ok = true
	if err != nil {
		log.Fatal(err)
		ok = false
	}
	return
}
