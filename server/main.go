package main

import (
	_ "aleshina/docs"
	"aleshina/repository"
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var repo repository.Repository

//	@title			Social Network API
//	@version		1.0
//	@description	API-документация для реализации простейшей соцсети

//	@contact.name	Ellina Aleshina
//	@contact.email	esalesina_1@edu.hse.ru

// @host		localhost:8080
// @BasePath	/api
func main() {
	// Оператор defer используется для отложенного вызова функции.
	// В данном случае - при завершении работы функции (и потока выполнения, то есть Горутины) Main
	defer log.Println("Shutting down completed")
	log.Println("Starting")

	log.Println("Opening database connection")
	// Инициализируем новый репозиторий
	repo.NewRepository()
	// Отложено закрываем подключение к БД
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
		log.Println("Connection to database closed successfully")
	}(repo.Db)

	// Контексты используются для Gracefully Shutdown, то есть для мягкого завершения работы программы
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		// В отдельной горутине проверяем поступление в специальный канал сигнала os.Interrupt
		sigChannel := make(chan os.Signal, 1)
		// В случае получения сигнала ^C(Interrupt) выполняем системный вызов на завершение работы программы
		signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)
		<-sigChannel
		close(sigChannel)
		// После получения сигнала завершаем выполнение функции
		cancel()
	}()
	// В основном потоке запускаем сервер
	Serve(ctx)
}
