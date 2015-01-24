package models

import (
	"fmt"
	"testing"
)

func TestSnippetStruct(t *testing.T) {
	snippet := Snippet{LanguageCode: "python",
		SessionId:   "12erfrvbx",
		Title:       "A simple python script",
		Description: "Adds two numbers",
		Tags:        []string{"simple", "python", "math"},
		Code:        "1 + 2",
		Deps:        []string{}}
	fmt.Printf("%+v\n", snippet)
}

func TestCreateSnippet(t *testing.T) {
	snippet := Snippet{LanguageCode: "python",
		SessionId:   "12erfrvbx",
		Title:       "A simple python script",
		Description: "Adds two numbers",
		Tags:        []string{"simple", "python", "math"},
		Code:        "1 + 2",
		Deps:        []string{}}
	snippetId, ok, _ := CreateSnippet(&snippet)
	fmt.Printf("%s\n", snippetId)
	if ok != true {
		t.Error("Expected", true, "but got", ok, "instead")
	}
}

func TestFindSnippetById(t *testing.T) {
	snippetId := "54c3b0cb900d19bf71d013d3"
	_, err := FindSnippetById(snippetId)
	if err != nil {
		t.Error("Failed querying the db")
	}
}

func TestUpdateSnippetById(t *testing.T) {
	snippetId := "54c3f73f1233bd6011000001"
	snippet := Snippet{LanguageCode: "ruby",
		SessionId:   "12erfrvbx",
		Title:       "A simple ruby script",
		Description: "Adds two numbers",
		Tags:        []string{"simple", "ruby", "math"},
		Code:        "1 + 2",
		Deps:        []string{}}
	ok, err := UpdateSnippetById(snippetId, &snippet)
	if err != nil {
		t.Error("Expected", true, "but got", ok, "instead")
	}
}

func TestFindSnippetsByUser(t *testing.T) {
	snippetUser := "12erfrvbx"
	_, err := FindSnippetsByUser(snippetUser)
	if err != nil {
		t.Error("Failed querying the db")
	}
}

func TestFindSnippetsByTag(t *testing.T) {
	snippetTag := "simple"
	snippets, err := FindSnippetsByTag(snippetTag)
	if err != nil {
		t.Error("Failed querying the db")
	}
	for _, snippet := range snippets {
		status := false
		for _, tag := range snippet.Tags {
			if tag == snippetTag {
				status = true
				break
			}
		}
		if status != true {
			t.Error("Snippet %+v\n", snippet, "doesn't contain tag %s", snippetTag)
		}
	}
}

//func TestDeleteSnippetById(t *testing.T) {
//	snippetId := "54c3f8771233bd60d3000001"
//	ok, _ := DeleteSnippetById(snippetId)
//	if ok != true {
//		t.Error("Failed querying the db")
//	}
//}
