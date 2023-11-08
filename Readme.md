Генерация grpc обёртки и swagger файла производилась с помощью утилиты buf (команда указана в Makefile).

Миграция для создания таблицы в базе данных производилась с помощью goose командой goose -dir migrations postgres "host=127.0.0.1 port=5432  user=el password=ozon sslmode=disable dbname=db" up.

Для подключения к базе можно использовать команду docker-compose up.
Запуск сервиса произваодится командой go run cmd/main.go {флаг}, где флагом может быть -db для хранения ссылок в базе данных и -im для хранения in-memory. 
Аналогично это нужно делать в docker-контейнере, забилдив Dockerfile.

Раз в сутки база или кэш чистятся с помощью кроны, ttl также равен одним суткам.

Моки для тестов генерировались с помощью утилиты gomock.

Для тестирования сервиса удобнее всего импортировать swagger файл в Postman, формат запросов указан в proto/url_shorter.proto.Сервис слушает на порту 8081.

Генерация укороченной ссылки производится в цикле до тех пор, пока полученная ссылка не будет уникальной.
