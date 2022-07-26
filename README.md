# Build Run Local
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
## Use air autoload 
```
docker-compose exec crawl air -c .air.conf
```

# Deploy
## Run file build.sh build project into folder bin
```
chmod +x build.sh
./build.sh
```

# Create Services in run in background (https://www.atpeaz.com/running-go-app-as-a-service-on-ubuntu/amp/)
## Create Service and Run App Web
```
sudo nano /lib/systemd/system/app_web.service
```
## Copy Content
```
[Unit]
Description=App Web

[Service]
Type=simple
Restart=always
RestartSec=5s
WorkingDirectory=/root/actions-runner/crawl/crawl/crawl/bin
ExecStart=/root/actions-runner/crawl/crawl/crawl/bin/app_web

[Install]
WantedBy=multi-user.target
```
```
sudo systemctl enable app_web
sudo systemctl start app_web
sudo systemctl status app_web
```

## Run App Crawl
```
./app_crawl
```

## Add CronTab
```
crontab -e
```
### add cron time
```
*/60 * * * * /root/actions-runner/crawl/crawl/crawl/bin/app_crawl
```
### Reload cron run
```
sudo service cron reload
```

## Website 
http://techdaily.info/