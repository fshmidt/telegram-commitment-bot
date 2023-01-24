# Тестовое задание avitoTech

<!-- ToC start -->
# Содержание

1. [Описание задачи](#Описание-задачи)
1. [Реализация](#Реализация)
1. [Endpoints](#Endpoints)
1. [Запуск](#Запуск)
1. [Примеры](#Примеры)
<!-- ToC end -->

# Описание бота

Этот бот контроллирует ваши обещания, запоминает дедлайны и напоминает о них. От обычного календаря он отличается тем, что находясь с вами в чатах с общими знакомыми, товарищами и друзьями, он будет стыдить и напоминать о приближающемся дедлайне, делая таким образом невыполнение обещания более неприятным.
# Реализация

- Подход "Чистой Архитектуры" и техника внедрения зависимости.
- Работа с NoSQL базой данных [bolt](https://github.com/boltdb/bolt).
- Конфигурация приложения - библиотека [viper](https://github.com/spf13/viper).
- Запуск из Docker.
- CI/CD Github-DigitalOcean
  **Структура проекта:**
```
.
├── pkg
│   ├── config      // парсинг конфигурации
│   ├── telegram    // логика бота
│   └── repository  // взаимодействие с БД
├── cmd             // точка входа в приложение
```

# Адрес бота

https://t.me/obeshyalkin_bot

# Запуск

```
make build-image
make start-container
```
Если приложение запускается впервые, необходимо прокинуть ваш токен телеграма в окружение.

# Примеры

### 1. /start

![start-obeshyalkin.png](/docs/start-obeshyalkin.png)

### 2. /promise

![promise-obeshyalkin.png](/docs/promise-obeshyalkin.png)

### 4. /mypromises

![mypromises-obeshyalkin.png](/docs/mypromises-obeshyalkin.png)

### 5. /delete

![delete-obeshyalkin.png](/docs/delete-obeshyalkin.png)

### 5. Напоминания

![remider-obeshyalkin-1.png](/docs/remider-obeshyalkin-1.png)
![remider-obeshyalkin-2.png](/docs/remider-obeshyalkin-2.png)
![remider-obeshyalkin-3.png](/docs/remider-obeshyalkin-3.png)

### 6. Реакция на фразы-триггеры

![proposition-obeshyalkin.png](/docs/proposition-obeshyalkin.png)


