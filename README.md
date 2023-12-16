# Web Apps Task 5
Приложение — простой аналог социальной сети с двумя сущностями - пользователи и публикации.  
Для разработки приложения были использованы следующие технологии:
- Язык Go
  + Фреймворк `Gorilla/MUX` для удобного роутинга
  + Собственная библиотека `net/http` для обработки запросов
  + Собственная библиотека `database/sql` для работы с SQL запросами
  + Собственная библиотека `html/template` для обработки HTML шаблонов
  + Собственные средства go `go vet` и `go lint` в GitHub Actions
- Язык SQL
- Язык HTML
- Docker и Docker Compose

Для запуска приложения требуется клонирование исходного репозитория и несколько команд в терминале Bash:  

```shell
cd <путь к директории проекта>
docker-compose build
docker-compose run -d
```

При первом запуске на сервере (или на локальной машине) требуется произвести миграцию базы данных с использованием [Golang Migrate](https://github.com/golang-migrate/migrate):

```shell
cd <путь к директории проекта>
go install -v github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate -path ./schema -database 'postgres://<username>:<password>@<host>:<port>/<db_name>?sslmode=<disable/enable>' up
```

Для доступа к страницам приложения достаточно перейти по ссылке http://localhost:8080/  
А также для доступа к Swagger-документации, требуется перейти по ссылке http://localhost:8080/swagger/index.html#

© Ellina Aleshina
