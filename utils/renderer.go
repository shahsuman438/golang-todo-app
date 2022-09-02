package utils

import (
	"fmt"
	"net/http"
	"text/template"
)

type Todo struct {
	Title      string
	IsComplete bool
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
	}
}
