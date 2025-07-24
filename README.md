# Тестовое задание для Customapp

## Запуск через Go

1. Убедитесь, что установлен Go версии 1.24.5 или выше.
2. Откройте терминал в папке проекта.
3. Запустите команду:

```sh
go run main.go -rtp=1.0
```

Параметр `-rtp` — целевое значение RTP (от 0.0 до 1.0). Например:

```sh
go run main.go -rtp=0.95
```

Сервис будет доступен по адресу: [http://localhost:64333/get](http://localhost:64333/get)

## Сборка и запуск бинарника

```sh
go build -o multiplexer
./multiplexer -rtp=1.0
```

## Запуск через Docker

```sh
docker build -t multiplexer-service .
docker run -p 64333:64333 multiplexer-service
```
