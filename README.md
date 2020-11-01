## SMS Sender

Инструкция по настройке автоматической отправки sms для проекта Марш. \
На выполнение этих шагов вы потратите 15-30 минут. \
В итоге вы абсолютно бесплатно и анонимно сможете получать самую актуальную информацию о точках на карте от проекта Марш при полном отсутствии интернета.

## Простой способ

#### Twilio(free):
Сервис отправки смс. \
Зарегистрируйтесь: https://www.twilio.com/try-twilio \
Используйте номер, на который хотите получать смс. Если у вас есть иностранная симка - используйте её. \
После регистрации у вас будет trial аккаунт, его хватит на 3+ маршей. \
Заведите Trial number.
Информация, которая вам нужна будет дальше: [Screenshot](https://raw.githubusercontent.com/InternetLess/sender/master/screenshots/twilio.png)
- Tiral Number
- Account Sid
- Auth Token

#### Google Cloud(free):
Облачный сервис, после регистрации у вас будет 300$ trial. Этого с лихвой хватит, чтобы запустить наш контейнер.
1. https://console.cloud.google.com/ Можете войти, используя существующий аккаунт google.
2. Активируйте trial. [Screenshot](https://raw.githubusercontent.com/InternetLess/sender/master/screenshots/activate-gcp-trial.png). Там потребуется карта, но это бесплатно и гуглу можно доверять :)
3. Зайдите на страницу создания сервиса Cloud Run(это дешёвый способ запускать контейнеры). https://console.cloud.google.com/run/create
4. Заполните поля, как на скриншоте. [Screenshot](https://raw.githubusercontent.com/InternetLess/sender/master/screenshots/cloud-run-fill-1.png)
5. Url образа контейнера: `gcr.io/internetless-sender/sender:latest` [Screenshot](https://raw.githubusercontent.com/InternetLess/sender/master/screenshots/cloud-run-fill-2.png)
6. Заполните переменные: [Screenshot](https://raw.githubusercontent.com/InternetLess/sender/master/screenshots/cloud-run-fill-3.png)
    - TWILIO_SID - получили при регистрации на twilio (Account Sid)
    - TWILIO_TOKEN - тоже при регистрации на twilio (Auth Token)
    - FROM_NUMBER - номер с twilio. Американский +1XXX
    - TO_NUMBER - ваш номер телефона, на который вы регистрировали twilio
7. Проверяете, что всё успешно. Копируете ссылку [Screenshot](https://raw.githubusercontent.com/InternetLess/sender/master/screenshots/cloud-run-result.png)

#### Ссылку нам:
Пришлите полученную ссылку в endpoints боту `@meshmarchbot` в телеграм. \
Ура! Настройка прошла успешно. Прелести этого подхода:
- Бесплатная отправка важных данных по смс, в случае когда нет интернета
- Анонимно. Даже создатели системы не знают ваш номер телефона

## Опционально, для технарей
Вы можете:
- Проверить код: `main.go` в репозитории, чтобы убедиться, что он никуда не посылает ваши данные.
- Собрать контейнер сомостоятельно и убедиться, что он такой же, как опубликованный: `gcr.io/internetless/sender:latest` \
 `docker build -t sender:latest .` \
  No dependencies except docker itself.
- Задеплоить этот контейнер любым удобным способом и скинуть endpoint нам: \
`docker run -e PORT=7000 -e FROM_NUMBER="+1XXX" -e TO_NUMBER="+1XXX" -e TWILIO_SID="XXX" -e TWILIO_TOKEN="XXX" -p 7000:7000 sender:latest` \
All env-vars are required except PORT. Default port is 8080.
- Использовать docker-compose для более удобной сборки (не забудьте изменить файл конфигурации!)
Для этого ```docker-compose up --build```
