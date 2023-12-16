package repository

import (
	"aleshina/initializr"
	"database/sql"
	"log"
)

// Repository - структура, содержащая в себе открытое подключение к БД
type Repository struct {
	Db *sql.DB
}

// NewRepository инициализирует репозиторй, открывая подключение к базе данных
func (r *Repository) NewRepository() {
	// Вызываем метод инициализатора подключения к БД
	db, err := initializr.DbConnectionInit()
	if err != nil {
		log.Fatal(err)
	}
	r.Db = db
}
