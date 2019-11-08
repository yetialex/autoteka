autoteka
======

Web сервис предоставляет REST API для управления сущностями типа Автомобиль. Сущности хранятся в MySQL.

Для сборки сервиса необходимы make и docker.

Сборка осуществляется в отдельном контейнере.

Postman коллекция в папке postman

## Быстрый старт

    make build
    docker-compose build
    docker-compose up

## Запуск тестов

    make test
    
## Примеры запросов

    Запрос на создание сущности:
        curl -H Content-Type:application/json -H Accept:application/json -X POST -d '{"id":1,"brand":"skoda","model":"yeti","engine_volume":1.2}' http://127.0.0.1:8092/autos
    Ответ:
        {"success":true,"message":"created auto with id: 1","payload":null}
    
    
    Запрос на чтение сущности:
            curl -X GET http://127.0.0.1:8092/autos/1
    Ответ:
        {"id":1,"brand":"skoda","model":"yeti","engine_volume":1.2}
        
        
    Запрос на обновление сущности:
            curl -H Content-Type:application/json -H Accept:application/json -X PUT -d '{"id":1,"brand":"skoda","model":"kodiaq","engine_volume":2.0}' http://127.0.0.1:8092/autos
    Ответ:
            {"success":true,"message":"updated auto with id: 1","payload":null}
            
          
    Запрос на удаление сущности:
        curl -X DELETE http://127.0.0.1:8092/autos/1
    
        {"success":true,"message":"deleted auto with id: 1","payload":null}