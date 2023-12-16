package main

import (
	"context"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log"
	"net/http"
)

// Serve - является функцией работы сервера. В ней указываются все параметры и она по сути постоянно запущена
func Serve(ctx context.Context) {
	server := http.Server{Addr: ":8080"}
	// Роутер используется для указания маршрутов на сервере
	router := mux.NewRouter()
	http.Handle("/", router)
	// Метод для подключения Swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	// Роутинг по шаблонам для отображения
	router.HandleFunc("/", HomeHandlerTmpl)

	// Пути к страницам, где отрисовываются HTML шаблоны с постами
	router.HandleFunc("/posts", FeedHandlerTmpl)
	router.HandleFunc("/posts/newpost", NewPostHandlerTmpl)
	router.HandleFunc("/posts/{id}", PostHandlerTmpl)
	router.HandleFunc("/posts/{id}/edit", EditPostHandlerTmpl)

	// Пути к страницам, где отрисовываются HTML шаблоны с пользователями
	router.HandleFunc("/users", UsersHandlerTmpl)
	router.HandleFunc("/users/register", RegisterHandlerTmpl)
	router.HandleFunc("/users/{id}", ProfileHandlerTmpl)
	router.HandleFunc("/users/{id}/edit", EditProfileHandlerTmpl)

	// Пути к API, то есть нет отрисовки шаблонов, данные передаются через JSON
	// Пути к API методам пользователя
	router.HandleFunc("/api/users/register", RegisterHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/users/{id}/edit", EditProfileHandler).Methods(http.MethodPut)
	router.HandleFunc("/api/users/{id}/delete", DeleteProfileHandler).Methods(http.MethodDelete)
	router.HandleFunc("/api/users", UsersHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/users/{id}", ProfileHandler).Methods(http.MethodGet)

	// Пути к API методам постов
	router.HandleFunc("/api/posts/newpost", NewPostHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/posts/{id}/edit", EditPostHandler).Methods(http.MethodPut)
	router.HandleFunc("/api/posts/{id}/delete", DeletePostHandler).Methods(http.MethodDelete)
	router.HandleFunc("/api/posts", FeedHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/posts/{id}", PostHandler).Methods(http.MethodGet)

	// В отдельном потоке запускаем сервер
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			return
		}
	}()
	// В основном потоке проверяем сигнал OS.INTERRUPT
	for {
		select {
		case <-ctx.Done():
			// Если получили сигнал об остановке, то завершаем работу сервера
			log.Println("Shutting down server")
			err := server.Shutdown(ctx)
			if err != nil {
				panic(err)
			}
			return
		}

	}
}
