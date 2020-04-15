# vscale-task

### Запуск программы

```shell script
go run main.go --port 8181 --token zz9740mdsk889732n34gnk20
```

### Создание группы серверов

В качестве параметра передается коичество необходимых копий, а в теле запроса передается требуемая конфигурация
```shell script
curl -X POST 'localhost:8181/7' -H 'Content-Type: application/json;charset=UTF-8' -d '{"make_from":"ubuntu_18.04_64_001_master","rplan":"small","do_start":false,"name":"New-Test","keys":[16],"location":"spb0"}'
```

### Удаление группы серверов

В качестве параметра передается идентификатор группы серверов

```shell script
curl 'localhost:8181/32' -X DELETE  
```

### Получение статуса группы серверов

В качестве параметра передается идентификатор группы серверов

```shell script
curl 'localhost:8181/32'
```
