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

### Build swagger API document

```sh
./scripts/swagger-build.sh
```

### Start server

```
go run main.go
```
