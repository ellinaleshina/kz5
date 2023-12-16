package models

// User представляет собой сущность
// пользователя социальной сети. У каждого пользователя
// есть уникальный id и никнейм.
type User struct {
	ID       int
	Username string
}
