# Тестовое задание для поступления в DevOps Cloud.ru Camp
Результаты выполнения тестового задания следует опубликовать на GitHub или захостить на любой открытой платформе (например, Github Pages) и отправить на почту devopscloudcamp@cloud.ru. Также следует указать свои контактные данные для получения обратной связи.

Тестовое задание состоит из двух задач. 
## 1. Ansible playbook
Необходимо написать Ansible playbook, который выполняет на хосте следующие действия:

- создает нового пользователя cloudru с паролем cloudpass 
- разрешает на хосте авторизацию через ssh по ключу
- запрещает логин по ssh от пользователя root 
- копирует предоставленный публичный ключ для пользователя cloudru

Плейбук должен выполняться относительно чистого дистрибутива ОС Ubuntu Server 22.04.3. Для написания и проверки плейбука можно локально развернуть VM с помощью любого удобного инструмента виртуализации (VirtualBox, VmWare Fusion, VmWare Fusion Player, Hyper-V, ...)

Публичный ключ для публикации на хост:
```
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCfrfE0OluoNHb5dOpV4RpWmVXvMBWc17kaM7DDjCm7romNQMDX95i5Fc67Q3c47pvrm/qi/ZqsCeqNdLl5+VV41rVz701Pj/UUr2FZpIm80Ur0iM1DFy81GKo/lS1INopqdd4KvUnM2d8yqfJSm9m5Cq7AM9S0mqObuMayfqNR4YcOlm9fnEMqhrSWbBVvdghPNiBzs7T9RzEq/0w8rs743tCF7MICv72fdgYadrGlxFsFWSujwZXQLI4VUSxKirJBCUgfR0u84gZK/wUzJ4EPqMichniTf24AsvidozUHWMDmQ+pUaBTyxjD5egi8LcV0EHH4feHwzacA2gyGbOtFK3wpa/dgE1yvPTkPKnccIXKnbel0mfxfsBVkclc5/DnczmrdaGrX5DCrQbI+HO4lhr4KzAm/pw6qfLcw8KjCdVKsnCRXykdat8KUwNAeolknRWdKDqdsbyXBj+ePMTlMR8YmoBj9znYWwOnAAyu56utiteL0oq9YPkb7ZGF5ZOE=
```
Полученный плейбук и команду для его запуска положить в папку /playbook

## 2. Web приложение на Python
### Приложение
Требуется написать простое веб-приложение на Python, которое слушает входящие соединения на порту 8000 и предоставляет HTTP API, в котором реализовано 3 метода:
- GET /hostname - при запросе на этот метод приложение отдает имя хоста, на котором запущено приложение
- GET /author - возвращает значение переменной окружения $AUTHOR, в которой задано имя или никнейм человека, выполняющего это задание
- GET /id - возвращает значение переменной окружения $UUID, содержащее любую произвольную строку-идентификатор в формате uuid

### Dockerfile
Необходимо написать Dockerfile для полученного приложения в соответствии с принятыми в сообществе best-practice.

Полученный скрипт и Dockerfile к нему положить в папку /app

### Kubernetes manifest
Далее необходимо написать манифест для запуска приложения в Kubernetes в отдельном неймспейсе в виде Deployment с 3 репликами и сервиса с типом ClusterIP. Реализовать readiness- и liveness- пробы. В переменную UUID должен подставляться уникальный идентификатор пода в кластере, в котором запущено приложение.

Манифест положить в папку /manifest

Для локального запуска кластера можно использовать инструменты Docker Desktop, kind, minikube и другие 

### Helm chart
Написать Helm чарт, в котором через переменные в файле values.yaml можно задать:
- имя образа, запускаемого в поде 
- количество реплик приложения
- значение, подставляемое в переменную AUTHOR

Полученный чарт положить в папку /helm
