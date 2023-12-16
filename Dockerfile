# Устанавливаем язык Go в контейнер
FROM golang:1.21

# Проверяем, что язык установлен
RUN go version
ENV GOPATH=/

# Копируем файлы проекта в контейнер в корень GOPATH
COPY ./ ./

# Устанавливаем postgresql
RUN apt-get update
RUN apt-get -y install postgresql-client


# Собираем Go приложение
RUN go mod download
RUN go build -o backend ./server

# Запускаем приложение
CMD ["./backend"]
