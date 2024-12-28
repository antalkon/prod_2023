Установка postgres:
```
docker run -d \
  --name pg_prod_23 \
  -e POSTGRES_USER=admin \
  -e POSTGRES_PASSWORD=qwerty \
  -e POSTGRES_DB=prod \
  -p 5601:5432 \
  postgres
```