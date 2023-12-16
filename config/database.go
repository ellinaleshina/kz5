package config

// Конфигурация для подключения к базе данных
const (
	host     = "host=fullstack-postgres"
	port     = "port=5432"
	user     = "user=adm"
	password = "password=pwd"
	dbname   = "dbname=aleshina"
	sslmode  = "sslmode=disable"
)

// ConnStr - строка подключения к базе данных
const ConnStr = host + " " + port + " " + user + " " + password + " " + dbname + " " + sslmode
