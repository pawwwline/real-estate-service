# real-estate-service

**Тестовое задание для отбора на Avito Backend Bootcamp 2024**

## Инструкция по запуску

1. Клонируйте репозиторий:
    ```bash
    git clone https://github.com/pawwwline/real-estate-service.git
    ```

2. Запустите сервис:
    ```bash
    docker-compose up --build
    ```

## Использование

### Авторизация

Отправьте GET запрос на `/dummyLogin` для получения токенов.

- Запрос для пользователя:
    ```bash
    curl -X GET "http://localhost:8080/api/v1/dummyLogin?user_type=client"
    ```

- Запрос для модератора:
    ```bash
    curl -X GET "http://localhost:8080/api/v1/dummyLogin?user_type=moderator"
    ```

### Создание нового дома (только для модераторов)

Отправьте POST запрос на `/house/create` с JSON телом запроса.

Пример запроса:
```bash
curl -X POST http://localhost:8080/api/v1/house/create \
-H "Authorization: Bearer <moderator_token>" \
-H "Content-Type: application/json" \
-d '{"address": "123 Main St", "year": 2024, "developer": "ABC Realty"}'
```


### Получение списка всех квартир по номеру дома

Отправьте GET запрос на /house/{id}

Пример запроса:

```bash
curl -X GET "http://localhost:8080/api/v1/house/123456 \
-H "Authoriazation: Bearer <token>"
```



### Создание квартиры

Отправьте POST запрос на /flat/create

Пример запроса:

```bash
curl -X POST http://localhost:8080/api/v1/flat/create \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX3R5cGUiOiJ1c2VyIn0.QzcYKfysTUDKFVcLW8T3A_YODiLmXFyP2_CP2kKZeDY" \
-H "Content-Type: application/json" \
-d '{"house_id": 12345, "price": 10000, "rooms": 4}'
```



### Обновления статуса квартиры_(только для модераторов)
Отправьте POST запрос на /flat/update

Пример запроса:

```bash
curl -X POST http://localhost:8080/api/v1/flat/update \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX3R5cGUiOiJtb2RlcmF0b3IifQ.sy5Cgo6lkmgptgG4RggKA4Qwfregs472xP72gMX1upg" \
-H "Content-Type: application/json" \
-d '{"id": 1, "status": "approved"}'
```



