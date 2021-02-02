# Тестовое задание

Примеры запросов:  
    - Добавление: curl -X POST -H 'Content-Type: application/json' --data '{"name":"utest"}' localhost:8080/users  
    - Список всех юзеров: curl localhost:8080/users  
    - Получение по id: curl localhost:8080/users/3  
    - Редактирование: curl -X PUT -H 'Content-Type: application/json' --data '{"name":"utest123"}' localhost:8080/users/3  
    - Удаление: curl -X DELETE localhost:8080/users/4  
