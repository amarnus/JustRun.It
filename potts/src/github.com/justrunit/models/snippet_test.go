package models

import (
	"fmt"
	"testing"
)

func TestSnippetStruct(t *testing.T) {
	snippet := Snippet{"python",
		"12erfrvbx",
		"A simple python script",
		"Adds two numbers",
		[]string{"simple", "python", "math"},
		"1 + 2",
		[]string{}}
	fmt.Printf("%+v\n", snippet)
}

func TestCreateSnippet(t *testing.T) {
	snippet := Snippet{"python",
		"12erfrvbx",
		"A simple python script",
		"Adds two numbers",
		[]string{"simple", "python", "math"},
		"1 + 2",
		[]string{}}
	ok, _ := CreateSnippet(&snippet)
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
	snippetId := "54c3b0cb900d19bf71d013d3"
	snippet := Snippet{"ruby",
		"12erfrvbx",
		"A simple ruby script",
		"Adds two numbers",
		[]string{"simple", "ruby", "math"},
		"1 + 2",
		[]string{}}
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

/*
func TestDeleteSnippetById(t *testing.T) {
	snippetId := "54c3e330900d19bf71d013e3"
	ok, _ := DeleteSnippetById(snippetId)
	if ok == true {
		t.Error("Failed querying the db")
	}
}
*/
