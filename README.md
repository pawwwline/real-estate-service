# real-estate-service

**Тестовое задание для отбора на Avito Backend Bootcamp 2024**

**Инструкция по запуску:**


1.Клонировать репозиторий
```git clone https://github.com/pawwwline/real-estate-service.git```


2.Запуск:
```docker-compose up --build``` 

**Использование**

**_Авторизация_**

Отправьте GET запрос на /dummyLogin

Запрос для пользователя:

```curl -X GET "http://localhost:8080/api/v1/dummyLogin?user_type=client"```

Запрос для модератора:

```curl -X GET "http://localhost:8080/api/v1dummyLogin?user_type=moderator"```



**_Создание нового дома (только для модераторов)_**

Отправьте POST запрос на /houses с JSON телом запроса. 

Пример запроса

```curl -X POST http://localhost:8080/api/v1/house/create \
-H "Authorization: Bearer <moderator_token>" \
-H "Content-Type: application/json" \
-d '{"address": "123 Main St", "year": 2024, "developer": "ABC Realty"}'```



**_Получение списка всех квартир по номеру дома_
**
Отправьте GET запрос на /house/{id}

Пример запроса

```curl -X GET "http://localhost:8080/api/v1/house/123456 -H "Authoriazation: Bearer <token>"```



**_Создание квартиры_**

Отправьте POST запрос на /flat/create

Пример запроса

``` curl -X POST http://localhost:8080/api/v1/flat/create \
 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX3R5cGUiOiJ1c2VyIn0.QzcYKfysTUDKFVcLW8T3A_YODiLmXFyP2_CP2kKZeDY" \
 -H "Content-Type: application/json" \
 -d '{"house_id": 12345, "price": 10000, "rooms": 4}'```



**_Обновления статуса квартиры_(только для модераторов)_**
Отправьте POST запрос на /flat/update

Пример запроса 

```curl -X POST http://localhost:8080/api/v1/flat/update \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX3R5cGUiOiJtb2RlcmF0b3IifQ.sy5Cgo6lkmgptgG4RggKA4Qwfregs472xP72gMX1upg" \
-H "Content-Type: application/json" \
-d '{"id": 1, "status": "approved"}'
```



