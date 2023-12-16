package main

import (
	"aleshina/models"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

// Ручки для отрисовки простейшего фронтенда с использованием шаблонов

// FeedHandlerTmpl - отображает ленту со всеми постами в шаблоне
func FeedHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	posts := repo.GetAllPosts()
	tmpl, _ := template.ParseFiles("templates/feed.html")
	type ViewData struct {
		Title string
		Posts []models.Post
	}
	data := ViewData{
		Title: "Feed",
		Posts: posts,
	}
	if err := tmpl.Execute(w, data); err != nil {
		return
	}
}

// NewPostHandlerTmpl - ручка создания нового поста в шаблоне
func NewPostHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/newPost.html")
	type ViewData struct {
		Title string
	}
	data := ViewData{
		Title: "New post",
	}
	if err := tmpl.Execute(w, data); err != nil {
		return
	}
}

// PostHandlerTmpl - ручка отображения поста по его ID в шаблоне
func PostHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/post.html")
	if err != nil {
		fmt.Print(err)
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	type ViewData struct {
		Post models.Post
	}
	data := ViewData{
		Post: repo.GetPostByID(id),
	}
	if err := tmpl.Execute(w, data); err != nil {
		return
	}
}

// EditPostHandlerTmpl - ручка шаблона изменения поста
func EditPostHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/editPost.html")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	post := repo.GetPostByID(id)
	type ViewData struct {
		Title string
		Post  models.Post
	}
	data := ViewData{
		Title: "Edit Post",
		Post:  post,
	}
	if err := tmpl.Execute(w, data); err != nil {
		return
	}
}
