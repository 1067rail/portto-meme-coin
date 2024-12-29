# Meme Coin

## Develop guide

### Start test db

```sh
docker volume create pg-data
docker run --name test-postgres \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=pass \
  -e POSTGRES_DB=meme-coin \
  -v pg-data:/var/lib/postgresql/data \
  -p 5432:5432 \
  -d postgres
```

### Create migration file

```sh
./scripts/migration-generate.sh
```

### Apply migration

```sh
./scripts/migration-apply.sh
```

### Start server

```
go run main.go
```

### Test API with cURL

1. create coin

```sh
curl -X POST \
-H "Content-Type: application/json" \
-d '{"name": "pepe", "description": "Pepe the Frog is a famous comic character and Internet meme created by cartoonist Matt Furie."}' \
http://localhost:8080/api/v1/coins 
```
