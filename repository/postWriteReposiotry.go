package repository

import (
	"aleshina/models"
	"fmt"
	"log"
)

// CreatePost - метод записи поста в базу данных. Для работы использует открытое подключение из Repository
func (r *Repository) CreatePost(post models.Post) (int, error) {
	var id int
	query := "INSERT INTO posts (author, posted, post_text) VALUES ($1, now(), $2) RETURNING id"

	// db.Query выполняет SQL запрос, который указан в строке выше, при этом может возвращать значение.
	// В качестве плейсхолдеров ($1, $2) используются значения post.Author.ID и post.PostText
	row, err := r.Db.Query(query, post.Author.ID, post.PostText)
	if err != nil {
		fmt.Println("key")
		log.Println(err)
		return 0, err
	}
	// Так как запрос возвращает ID, требуется его записать в переменную
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

// UpdatePost - метод изменения поста. На вход принимает пост, редактируя его в БД
func (r *Repository) UpdatePost(post models.Post) {
	query := "UPDATE posts SET post_text = $1 WHERE id = $2"
	// db.Exec выполняет SQL запрос без возврата значений
	_, err := r.Db.Exec(query, post.PostText, post.ID)
	if err != nil {
		return
	}
}

// DeletePost - метод удаления поста по его ID
func (r *Repository) DeletePost(id int) {
	query := "DELETE FROM posts WHERE id = $1"
	_, err := r.Db.Exec(query, id)
	if err != nil {
		return
	}
}
