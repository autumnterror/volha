## Серверная часть для кампании Volha

_development by [breezy innovations rzn](contacts.breezynotes.ru)_

Используется подход построения архитектуры gRPC + HTTP Gateway. 

Стек:
* Go (echo framework)
* PostgreSQL для хранения основной информации
* Redis для кеширования
* Docker для развертывания

Работа с товарами представлена в микросервисе product-service

Работа с кешрованием происходит в Gateway. Документация последнего описана через go-swag и храниться в директории gateway/docs/
