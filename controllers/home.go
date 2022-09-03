package controllers

import (
	"net/http"
	"todoapp/utils"
)

type Todo struct {
	Title      string
	IsComplete bool
}

func Home(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "Todo-home.tmpl")
}
