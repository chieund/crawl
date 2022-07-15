#chmod +x build.sh

## copy env to folder bin
cp app.yaml bin/

## build app web
go build -o bin/app_web

## build crawl app
go build -o ./bin/app_crawl ./cmd/