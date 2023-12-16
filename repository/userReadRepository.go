package repository

import (
	"aleshina/models"
	"log"
)

// GetAllUsers является методом получения всех пользователей социальной сети.
// Может использоваться администраторами соцсети, например для мониторинга
// активности пользователей.
func (r *Repository) GetAllUsers() []models.User {
	var users []models.User
	rows, err := r.Db.Query("select * from users ORDER BY id")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.ID, &u.Username)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, u)
	}
	err = rows.Close()
	if err != nil {
		log.Println(err)
		return nil
	}
	return users
}

// GetUserByID является методом получения пользователя по уникальному id.
// Может использоваться например для получения собственного профиля или
// для поиска конкретного пользователя.
func (r *Repository) GetUserByID(id int) models.User {
	var user models.User
	query := "SELECT * FROM users WHERE id = $1"
	row := r.Db.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Username)
	if err != nil {
		return models.User{}
	}
	return user
}
