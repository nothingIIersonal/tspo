# Технологии создания программного обеспечения
## Задание 666. Мы дошли до финала. Порадуемся этому, господа...

### Запуск без Docker
```
swag init --parseDependency --parseInternal  -d cmd/server/
go run cmd/server/main.go
```

URL: http://localhost:8080/

URL: http://localhost:8080/swagger/index.html

### Запуск в Docker
```
sudo docker compose up
```

URL: http://localhost:3000/

URL: http://localhost:3000/swagger/index.html
