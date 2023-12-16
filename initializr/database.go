package initializr

import (
	"aleshina/config"
	"database/sql"
	_ "github.com/lib/pq" // драйвер базы данных
	"log"
)

// DbConnectionInit инициализирует подключение к базе данных PostgreSQL
func DbConnectionInit() (*sql.DB, error) {
	// Явно указываем драйвер и строку с данными для подключения, после чего открываем подключение к БД
	db, err := sql.Open("postgres", config.ConnStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("Connection to database opened successfully")
	return db, nil
}
