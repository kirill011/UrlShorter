# UrlShorter
## Задание 
Требуется создать микросервис сокращения url. Длина сокращенного URL-адреса должна быть как
можно короче. Сокращенный URL может содержать цифры (0-9) и буквы (a-z, A-Z)
Эндпоинты:
### POST http://localhost:8080/
+ Request: (body): http://cjdr17afeihmk.biz/123/kdni9/z9d112423421
+ Response: http://localhost:8080/qtj5opu
### GET
+ Request (url query): http://localhost:8080/qtj5opu
+ Response (body): http://cjdr17afeihmk.biz/123/kdni9/z9d112423421

Микросервис должен уметь хранить информацию в памяти и в postgres в зависимости от флага
запуска -d

## Запуск
1. Отредактируйте параметры в файле configs/config.yaml
   + host - хост сервера базы данных
   + user - имя пользователя базы данных
   + password - пароль от пользователя БД
   + port - порт сервера базы данных
   + sslMode - SSLMode базы данных
   + timeZone - временная зона базы данных
   + dbName - название базы данных
   + runIp - хост, на котором будет запущен микросервис 
   + runPort - порт, на котором будет запущен сервис
3. Запустите 
  ``` cmd
  go run cmd/UrlShorter/main.go (-d)
  ```

## Запроы
+ POST запрос передаётся с Content-Type = application/json. В Body запроса должны находиться валидная ссылка. В ответ вернётся строка содержащая сокращённую ссылку
+ GET запрос должен быть следующего вида: http://{config.runIp}:{config.runPort}/{short_url}. В ответе вернётся строка full_url из базы(из памяти) которая соответствует short_url

Пример запросов находится в директории postmanExample
