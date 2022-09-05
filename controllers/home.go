package controllers

import (
	"net/http"
	"todoapp/utils"
)

type Todo struct {
	Title      string
	IsComplete bool
}

type heading struct {
	Title string
}

func Home(w http.ResponseWriter, r *http.Request) {
	p := heading{
		Title: "Golang Todo App",
	}
	utils.RenderTemplate(w, "Todo-home.tpl", p)
}
