package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Snippet struct {
	Id           bson.ObjectId `json:"_id"            bson:"_id"`
	LanguageCode string        `json:"language_code"  bson:"language_code"`
	SessionId    string        `json:"session_id"     bson:"session_id"`
	Title        string        `json:"title"          bson:"title"`
	Description  string        `json:"description"    bson:"description"`
	Tags         []string      `json:"tags"           bson:"tags"`
	Code         string        `json:"code"           bson:"code"`
	Deps         []string      `json:"deps"           bson:"deps"`
	IsPublic     bool          `json:"is_public"      bson:"is_public"`
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

func CreateSnippet(snippet *Snippet) (id string, ok bool, err error) {
	session, c, _ := getCollection("user_snippets", "snippets")
	defer session.Close()
	snippet.Id = bson.NewObjectId()
	id = snippet.Id.Hex()
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
	snippet.Id = bson.ObjectIdHex(snippetId)
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
	err = c.Find(bson.M{"session_id": snippetUser}).Iter().All(&snippets)
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
	if snippetTag != "" {
		err = c.Find(bson.M{"tags": snippetTag, "is_public": true}).Iter().All(&snippets)
	} else {
		err = c.Find(bson.M{"is_public": true}).Iter().All(&snippets)
	}
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
