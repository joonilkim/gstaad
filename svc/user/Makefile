root := ../..
ns := gstaad
service := user

test:
	go test

build-app:
	docker-compose build app

build-envoy:
	docker-compose build envoy

dev-app:
	docker-compose run --rm --service-ports app

dev-envoy:
	docker-compose run --rm --service-ports envoy

stop:
	docker-compose stop