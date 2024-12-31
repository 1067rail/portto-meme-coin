# Meme Coin

## Getting started

### Set up environment variables

```sh
cp example.env .env
# Edit the .env file to update configurations
```

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

2. get coin

```sh
curl -X GET \
http://localhost:8080/api/v1/coins/4681d911-c4f4-460d-918a-8c124d954626
```

3. update coin

```sh
curl -X PATCH \
-H "Content-Type: application/json" \
-d '{"name": "pepe", "description": "Pepe the Frog is a famous comic character and Internet meme created by cartoonist Matt Furie."}' \
http://localhost:8080/api/v1/coins/4681d911-c4f4-460d-918a-8c124d954626
```

4. delete coin

```sh
curl -X DELETE \
http://localhost:8080/api/v1/coins/4681d911-c4f4-460d-918a-8c124d954626
```

4. poke coin

```sh
curl -X POST \
http://localhost:8080/api/v1/coins/4681d911-c4f4-460d-918a-8c124d954626/poke
```

## Deployment guide

### Build docker image

```sh
docker build --tag meme-coin:latest .
```
