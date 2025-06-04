# Cache Based on Linked List

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Chi](https://img.shields.io/badge/chi-%23000000.svg?style=for-the-badge&logo=&logoColor=white)

---

## Описание проекта

Это учебный сервис на Golang с горутинами и многопоточностью с использованием роутера Chi, реализующий RESTful API с кэшированием на основе связного списка (LRU cache).

Сервис получает на вход массив чисел, возвращая результат их умножения на текущий множитель. При этом вычисления кэшируются с использованием LRU-кеша, чтобы ускорить повторные запросы и уменьшить задержки.

---

## Особенности и функциональные требования

- **Роутинг и API:**
```
GET  http://localhost:8080/info?numbers=2,789,7,59,neflefjl,fhgfg,eong5,egv
GET  http://localhost:8080/admin/cache/output
POST http://localhost:8080/admin/cache/changeSize
POST http://localhost:8080/admin/cache/changeFactor
POST http://localhost:8080/admin/worker/setActiveStatus
```

- **Кэширование (LRU):**

  - Используется связный список для обеспечения поиска и обновления элементов кэша.

  - Максимальный размер кэша задаётся переменной и может динамически изменяться через API.

- **Слой "Базы данных":**

  - "Запрос" к базе — умножение числа на текущий множитель (начинается с 4).

- **Множитель:**

  - Каждые 5 секунд множитель увеличивается на 1.

  - При каждом изменении множителя кэш необходимо пересчитать (перезаполнить), так как значения устаревают.

---

## Установка и запуск

1. Клонируйте репозиторий:

```
git clone https://github.com/yourusername/cache-linkedlist.git
```

2. Запустите программу

```
cd server_side
go build main.go
./main
```

3. Запустите http-клиент для выполнения запросов

```
cd client_side
go build main.go
./main
```
