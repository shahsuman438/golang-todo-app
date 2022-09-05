package controllers

import (
	"fmt"
	"net/http"
	"text/template"
)

type Todo struct {
	Title      string
	IsComplete bool
}

func Home(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("./templates/Todo-home.tpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
	}
}
