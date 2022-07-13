## Build Docker
```
docker-compose up --build
```

## Install package Golang
```
docker-compose exec crawl go mod tidy
```

## Folder vendor
```
docker-compose exec crawl go mod vendor
```

## Run Crawl
```
docker-compose exec crawl go run cmd/main.go
```