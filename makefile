project_build:
	go build -v ./...

project_test:
	go test -v ./...

copy_template:
	cp app.yaml bin/
	cp -R templates/ bin/

build_app_web:
	go build -o bin/app_web

build_app_crawl:
	go build -o ./bin/app_crawl ./cmd/

migrate:
	go build -o ./bin/app_migrate ./migrations/

restart_app_web:
	sudo systemctl restart app_web

reload_cronjob:
	sudo service cron reload

deploy_project: copy_template build_app_web build_app_crawl migrate restart_app_web reload_cronjob