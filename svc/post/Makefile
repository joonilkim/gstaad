root := ../..
ns := gstaad
service := post

test:
	go test

build:
	docker-compose build envoy app

build-envoy:
	docker-compose build envoy

dev: build-envoy
	docker-compose run --rm --service-ports envoy

stop:
	docker-compose stop