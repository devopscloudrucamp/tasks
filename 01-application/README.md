### Команды для сборки и загрузки образа в Docker Hub
- ``docker build -t tyumin000/devopscloudrucamp:echo-server-latest .``
- ``docker login``
- ``docker push tyumin000/devopscloudrucamp:echo-server-latest``
- 
### Запуск контейнера
- ``docker run -e AUTHOR=NIKITA -p 8000:8000 tyumin000/devopscloudrucamp:echo-server-latest``