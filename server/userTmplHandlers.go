package main

import (
	"aleshina/models"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Ручки для отрисовки простейшего фронтенда с использованием шаблонов

// HomeHandlerTmpl - ручка домашнего шаблона
func HomeHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	// Явно указываем локацию шаблона
	tmpl, _ := template.ParseFiles("templates/index.html")
	// Создаем структуру отображаемых данных
	type viewData struct {
		Title string
	}
	// Выполняем шаблон
	err := tmpl.Execute(w, viewData{Title: "WebApps Task 5"})
	if err != nil {
		log.Println(err)
	}

}

// RegisterHandlerTmpl - ручка шаблона с регистрацией
func RegisterHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/registration.html")
	type viewData struct {
		Title string
	}
	err := tmpl.Execute(w, viewData{Title: "New User"})
	if err != nil {
		log.Println(err)
	}
}

// ProfileHandlerTmpl - ручка шаблона с профилем пользователя
func ProfileHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/profile.html")
	// Получаем ID пользователя из адресной строки
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// Используя методы репозитория, получаем информацию о пользователе и его постах
	user := repo.GetUserByID(id)
	posts := repo.GetAllPostsByAuthor(id)
	type ViewData struct {
		Title string
		User  models.User
		Posts []models.Post
	}
	// Явно указываем отображаемую информацию
	data := ViewData{
		Title: "User profile",
		User:  user,
		Posts: posts,
	}
	if err := tmpl.Execute(w, data); err != nil {
		return
	}
}

// UsersHandlerTmpl - ручка для отображения шаблона со всеми пользователями
func UsersHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/users.html")
	users := repo.GetAllUsers()
	type ViewData struct {
		Title string
		Users []models.User
	}
	data := ViewData{
		Title: "Users List",
		Users: users,
	}
	if err := tmpl.Execute(w, data); err != nil {
		return
	}
}

// EditProfileHandlerTmpl - шаблон изменения профиля
func EditProfileHandlerTmpl(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/editProfile.html")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user := repo.GetUserByID(id)
	type ViewData struct {
		Title string
		User  models.User
	}
	data := ViewData{
		Title: "Edit Profile",
		User:  user,
	}
	if err := tmpl.Execute(w, data); err != nil {
		return
	}
}
