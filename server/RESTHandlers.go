package main

import (
	"aleshina/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// REST - ручки доступны по адресу http://<host>/api/<route>
// Ответом REST ручек являются JSON файлы
// Также для каждой ручки прописана Swagger документация

// RegisterHandler - ручка регистрации
//
//	@Summary		REST ручка для регистрации нового пользователя
//	@ID				register
//	@Accept mpfd
//	@Param		username	formData	string	true	"username"
//	@Produce json
//	@Tags users
//	@Router			/users/register [post]
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Для обеспечения работы в рамках localhost нужно явно указать допуск CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	// Парсим данные из формы регистраци
	err := r.ParseMultipartForm(10)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Получаем данные из полей
	data := r.Form.Get("username")
	var user models.User
	// Явно передаем их в структуру пользователя
	user.Username = data
	// Вызываем метод регистрации пользователя
	id, err := repo.RegisterUser(user)
	if err != nil {
		return
	}
	// Записываем ответ об успешном создании пользователя
	resp := []byte(fmt.Sprintf("User %d created successfully", id))
	_, err = w.Write(resp)
	if err != nil {
		return
	}
}

// EditProfileHandler - ручка изменения профиля
//
//	@Summary		REST - ручка для изменения профиля
//	@ID				edit-profile
//	@Accept mpfd
//	@Param		id			path		int		true	"userId"
//	@Param		username	formData	string	true	"username"
//	@Produce json
//	@Tags users
//	@Router			/users/{id}/edit [put]
func EditProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Change '*' to the specific origin you want to allow or use multiple origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Include the HTTP methods you want to allow
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := r.ParseMultipartForm(10)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}
	data := r.Form.Get("username")
	var user models.User
	user.Username = data
	user.ID = id
	repo.UpdateUser(user)
	if err != nil {
		return
	}
}

// DeleteProfileHandler - ручка удаления профиля
//
//	@Summary		REST ручка для удаления профиля
//	@ID				delete-profile
//	@Accept json
//	@Param		id	path		int		true	"User ID"
//
// @Produce json
// @Tags users
// @Router			/users/{id}/delete [delete]
func DeleteProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Change '*' to the specific origin you want to allow or use multiple origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Include the HTTP methods you want to allow
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	// Получаем ID пользователя из адресной строки
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	repo.DeleteUser(id)
	_, err := w.Write([]byte("User deleted successfully"))
	if err != nil {
		return
	}

}

// UsersHandler - ручка получения всех пользователей
//
//	@Summary		REST ручка вывода всех пользователей
//	@ID				all-users
//	@Accept json
//	@Produce json
//	@Tags users
//	@Router			/users [get]
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Change '*' to the specific origin you want to allow or use multiple origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Include the HTTP methods you want to allow
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	// Сериализуем полученные данные из БД, чтобы вывести их в Swagger
	res, _ := json.Marshal(repo.GetAllUsers())
	// Явно указываем тип возвращаемого значения
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// ProfileHandler - ручка получения профиля по ID
//
//	@Summary		REST ручка для получения пользователя по ID
//	@ID				user-by-id
//	@Accept json
//	@Param		id	path		int		true	"User ID"
//	@Produce json
//	@Tags users
//	@Router			/users/{id} [get]
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Change '*' to the specific origin you want to allow or use multiple origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Include the HTTP methods you want to allow
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	res, _ := json.Marshal(repo.GetUserByID(id))
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// NewPostHandler - ручка создания поста
//
//	@Summary		REST ручка для создания нового поста
//	@ID				newpost
//	@Accept mpfd
//	@Param		postText	formData	string	true	"postText"
//	@Param		author	formData	int	true	"author"
//	@Produce json
//	@Tags posts
//	@Router			/posts/newpost [post]
func NewPostHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Change '*' to the specific origin you want to allow or use multiple origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Include the HTTP methods you want to allow
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	err := r.ParseMultipartForm(10)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}
	data := r.Form.Get("postText")
	id, _ := strconv.Atoi(r.Form.Get("author"))
	var post models.Post
	post.PostText = data
	post.Author.ID = id
	postID, err := repo.CreatePost(post)
	if err != nil {
		// Если возникает ошибка, то выводим, что пользователя, который должен быть автором поста не существует
		fmt.Println(err)
		http.Error(w, "This user doesnt exist", http.StatusNotFound)
		return
	}
	resp := []byte("Post " + strconv.Itoa(postID) + "created successfully")
	_, err = w.Write(resp)
	if err != nil {
		return
	}
}

// FeedHandler - ручка вывода всех постов
//
//	@Summary		REST ручка для получения всех постов
//	@ID				all-posts
//	@Accept json
//	@Produce json
//	@Tags posts
//	@Router			/posts [get]
func FeedHandler(w http.ResponseWriter, r *http.Request) {
	posts := repo.GetAllPosts()
	res, _ := json.Marshal(posts)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// EditPostHandler - ручка изменения поста
//
//	@Summary		REST ручка для изменения поста
//	@ID				edit-post
//	@Accept mpfd
//	@Param		postText	formData	string	true	"postText"
//	@Param		id			path		int		true	"postId"
//	@Produce json
//	@Tags posts
//	@Router			/posts/{id}/edit [put]
func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Change '*' to the specific origin you want to allow or use multiple origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Include the HTTP methods you want to allow
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// Parse form data
	err := r.ParseMultipartForm(10)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Access form fields
	data := r.Form.Get("postText")
	var post models.Post
	post.ID = id
	post.PostText = data
	repo.UpdatePost(post)
	if err != nil {
		return
	}
}

// DeletePostHandler - ручка удаления поста
//
//	@Summary		REST ручка удаления поста
//	@ID				delete-post
//	@Accept json
//	@Param		id			path		int		true	"postId"
//	@Produce json
//	@Tags posts
//	@Router			/posts/{id}/delete [delete]
func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Change '*' to the specific origin you want to allow or use multiple origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Include the HTTP methods you want to allow
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	repo.DeletePost(id)
	_, err := w.Write([]byte("Post deleted successfully"))
	if err != nil {
		return
	}
}

// PostHandler - ручка получения поста по ID
//
//	@Summary		REST ручка получения поста по ID
//	@ID				get-post-by-id
//	@Accept json
//	@Param		id			path		int		true	"postId"
//	@Produce json
//	@Tags posts
//	@Router			/posts/{id} [get]
func PostHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Change '*' to the specific origin you want to allow or use multiple origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Include the HTTP methods you want to allow
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	res, _ := json.Marshal(repo.GetPostByID(id))
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
