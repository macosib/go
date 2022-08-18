#  <font color='red'>Зависимости:</font>
* Необходимо установить [Docker](https://docs.docker.com/install/linux/docker-ce/ubuntu/#install-docker-ce-1), и установить [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).
* Необходимо установить [Docker Compose](https://docs.docker.com/compose/install/#install-compose).

#  <font color='red'>Установка:</font>
* Скачать проект с репозитория.
* Собрать все образы.
    - `docker-compose build`
* Запустить контейнеры.
    - `docker-compose up -d`

После установки для тестовых запросов можно использовать файл requests.http:
```url
http://localhost:8080/api/v1/
```

#  <font color='red'>Завершение работы:</font>

Для остановки приложения необходимо выполнить:
- `docker-compose down`

