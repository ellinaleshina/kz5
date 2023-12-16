package models

import "time"

// Post является сущностью,
// которая представляет пост в социальной
// сети. У каждого поста может быть автор,
// время публикации и содержимое поста.
type Post struct {
	ID       int
	Author   User
	Posted   time.Time
	PostText string
}
