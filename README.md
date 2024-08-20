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

```curl -X GET "http://localhost:8080/dummyLogin?usertype=client"```

Запрос для модератора:

```curl -X GET "http://localhost:8080/dummyLogin?usertype=moderator"```



**__Создание нового дома_ (только для модераторов)_**

Отправьте POST запрос на /houses с JSON телом запроса. 

Пример запроса

```curl -X POST http://localhost:8080/houses -d '{"address": "123 Main St", "year": 2024, "developer": "ABC Realty"}' -H "Content-Type: application/json"```



**_Получение списка всех квартир по номеру дома_
**
Отправьте GET запрос на /house/{id}

Пример запроса

```curl -X GET "http://localhost:8080/house/123456"```



**_Создание квартиры_**

Отправьте POST запрос на /flat/create

Пример запроса

```curl -X POST http://localhost:8080/flat/create -d '{"house_id": 12345, "price": 10000, "rooms": 4}' -H "Content-Type: application/json"```



**_Обновления статуса квартиры_(только для модераторов)_**
Отправьте POST запрос на /flat/update

Пример запроса 

```curl -X POST http://localhost:8080/flat/update -d '{"id": 123456, "status": "approved"}' -H "Content-Type: application/json"```



