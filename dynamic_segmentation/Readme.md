# Сервис динамического сегментирования пользователей

Микросервис сервис, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)


# Начало работы

Для подготовки сервиса необходимо:

Запустить сервис можно с помощью команды `make dc`
1) `make migrate-up`
2) `make run`

Для запуска теста:
1. `make test` запуск всех тестов

## Примеры

### Создание пользователя <a name="create-user"></a>

Создание пользователя с указанным именем:
```curl
curl --location --request POST 'http://localhost:8080/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Ivan"
}'
```
Пример ответа:
```json
{
  "message": "User created successfully",
  "user_id": 1
}
```

### Создание сегмента <a name="create-seg"></a>

При создании сегмента реализована опция указания процента пользователей (из общего колличества), которые попадут в этот сегмент автоматически, а так же есть возможность установить TTL:
```curl
curl --location --request POST 'http://localhost:8080/segment' \
--header 'Content-Type: application/json' \
--data-raw '{
    "slug": "AVITO_SALE_30",
    "expiration_date": "2023-12-31T23:59:59Z",
    "random_percentage": 0.0
}'
```
Пример ответа:
```json
{
   "message": "Segment and user assignments created successfully"
}
```
### Удаление пользователя <a name="del-user"></a>

Удаление пользователя по указанному user_id:
```curl
curl --location --request DELETE 'http://localhost:8080/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": 1
}'
```
Пример ответа:
```json
{
   "message": "User deleted successfully",
   "user_id": 1
}
```

### Удаление сегмента <a name="del-seg"></a>

Удаление сегмента по указанному slug:
```curl
curl --location --request DELETE 'http://localhost:8080/segment' \
--header 'Content-Type: application/json' \
--data-raw '{
    "slug": "AVITO_SALE_30"
}'
```
Пример ответа:
```json
{
   "message": "Segment deleted successfully",
   "segment_id": 1
}
```

### Добавление/Удаление сегментов <a name="add-remove"></a>

Добавление / удаление сегментов пользователя списком без перетирания существующих сегментов с возможностью установить TTL.
```curl
curl --location --request POST 'http://localhost:8080/user/segments' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": 1,
    "add": [
    {
      "slug": "AVITO_SALE_10",
      "expiration_date": "2023-12-31T23:59:59Z"
    },
    {
      "slug": "AVITO_SALE_30",
      "expiration_date": "2023-11-30T23:59:59Z"
    },
    {
      "slug": "AVITO_SALE_20",
      "expiration_date": "2023-11-30T23:59:59Z"
    }
    ], 
    "remove": ["AVITO_SALE_40"]
}'
```
Пример ответа:
```json
{
   "message": "User segments updated successfully",
   "user_id": 1
}

```

### Получение списка сегментов <a name="seg-list"></a>

Получение списка сегментов пользователя по id:
```curl
curl --location --request GET 'http://localhost:8080/user/segments' \
--header 'Content-Type: application/json' \
--data-raw '{
   "user_id": 1
}'
```
Пример ответа:
```json
{
   "segments": ["AVITO_SALE_30","AVITO_SALE_10"],
   "user_id": 1
}
```

