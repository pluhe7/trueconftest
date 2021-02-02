# Тестовое задание

Примеры запросов:
    1. Добавление: curl -X POST -H 'Content-Type: application/json' --data '{"name":"utest"}' localhost:8080/users
    2. Список всех юзеров: curl localhost:8080/users
    3. Получение по id: curl localhost:8080/users/3
    4. Редактирование: curl -X PUT -H 'Content-Type: application/json' --data '{"name":"utest123"}' localhost:8080/users/3
    5. Удаление: curl -X DELETE localhost:8080/users/4
