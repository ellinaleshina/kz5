package repository

import (
	"aleshina/models"
	"fmt"
	"log"
)

// GetAllPosts - метод для получения всех постов в соцсети (т. н. Feed)
func (r *Repository) GetAllPosts() []models.Post {
	var posts []models.Post
	var post models.Post
	query := "SELECT p.id, p.posted, p.post_text, u.id, u.username FROM posts p JOIN users u on p.author = u.id ORDER BY p.id"
	rows, err := r.Db.Query(query)
	if err != nil {
		log.Println(err)
		return nil
	}
	// После выполнения запроса сканируем каждую строку и записываем значения в структуру.
	for rows.Next() {
		err = rows.Scan(&post.ID, &post.Posted, &post.PostText, &post.Author.ID, &post.Author.Username)
		if err != nil {
			log.Println(err)
			return nil
		}
		posts = append(posts, post)
	}
	return posts
}

// GetAllPostsByAuthor - метод получения всех постов конкретного автора. На вход принимает id автора
func (r *Repository) GetAllPostsByAuthor(id int) []models.Post {
	var posts []models.Post
	var post models.Post
	query := "SELECT p.id, p.posted, p.post_text, u.id, u.username FROM posts p JOIN users u on p.author = u.id where author = $1 ORDER BY p.id "
	rows, err := r.Db.Query(query, id)
	if err != nil {
		log.Println(err)
		return nil
	}
	for rows.Next() {
		err = rows.Scan(&post.ID, &post.Posted, &post.PostText, &post.Author.ID, &post.Author.Username)
		if err != nil {
			log.Println(err)
			return nil
		}
		posts = append(posts, post)
	}
	return posts
}

// GetPostByID - метод получения поста по его ID
func (r *Repository) GetPostByID(id int) models.Post {
	var post models.Post
	query := "SELECT p.id, p.posted, p.post_text, u.id, u.username FROM posts p JOIN users u on p.author = u.id where p.id = $1 ORDER BY p.id"
	err := r.Db.QueryRow(query, id).Scan(&post.ID, &post.Posted, &post.PostText, &post.Author.ID, &post.Author.Username)
	if err != nil {
		fmt.Print(err)
	}
	return post
}
