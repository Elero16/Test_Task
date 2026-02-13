# Subscription Aggregator Service (Test Task)

REST-сервис на Go для агрегации данных об онлайн-подписках пользователей.

## Стек технологий
- **Language:** Go 1.21
- **Framework:** Gin Gonic (HTTP Router)
- **Database:** PostgreSQL 15
- **Documentation:** Swagger (swag)
- **Containerization:** Docker & Docker Compose

## Функционал
- CRUD операции над подписками.
- Подсчет суммарной стоимости подписок пользователя за период (с фильтрацией по названию).
- Логирование всех операций.
- Автоматическое развертывание всей инфраструктуры через Docker.

## Как запустить

1. **Клонируйте репозиторий:**
   ```bash
   git clone [https://github.com/Elero16/effective-mobile-test.git](https://github.com/Elero16/effective-mobile-test.git)

   cd effective-mobile-test
