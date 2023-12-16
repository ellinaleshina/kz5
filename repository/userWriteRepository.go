package repository

import (
	"aleshina/models"
)

// RegisterUser является методом создания нового пользователя.
func (r *Repository) RegisterUser(user models.User) (int, error) {
	var id int
	query := "INSERT INTO users(username) VALUES ($1) RETURNING id"
	row := r.Db.QueryRow(query, user.Username)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateUser является методом изменения профиля пользователя
func (r *Repository) UpdateUser(user models.User) {
	query := "UPDATE users SET username = $1 WHERE id = $2"
	_, err := r.Db.Exec(query, user.Username, user.ID)
	if err != nil {
		return
	}
}

// DeleteUser удаляет профиль пользователя из базы данных
func (r *Repository) DeleteUser(id int) {
	query := "DELETE FROM posts WHERE author = $1"
	_, err := r.Db.Exec(query, id)
	if err != nil {
		return
	}
	query = "DELETE FROM users WHERE id = $1"
	_, err = r.Db.Exec(query, id)
	if err != nil {
		return
	}
}
