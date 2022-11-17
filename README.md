## Govorun

Для запуска вводим команду
```shell
make
```

Подписаться на сообщения curl
```shell
curl localhost:5000/listen
```

Изменить сообщение:
```shell
curl --request POST localhost:5000/say \
--header 'Content-Type: application/json' \
--data '{"word": "Hello World!!"}'
```