package template

import (
	"encoding/json"
	"github.com/justrunit/models"
	"os"
)

func GetTemplateByLanguage(language string) (snippet *models.Snippet, err error) {
	file, _ := os.Open("src/github.com/justrunit/template/template.json")
	decoder := json.NewDecoder(file)
	template := make(map[string]*models.Snippet)
	err = decoder.Decode(&template)
	snippet = template[language]
	return snippet, err
}
