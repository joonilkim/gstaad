root := ../..
ns := gstaad
service := gateway

test:
	go test

docker-build:
	docker build \
		-f Dockerfile \
		-t ${ns}:${service} ${root}

dev: docker-build
	docker run \
		-p 8080:8080 \
		--env-file .env --rm \
		--name ${ns}-${service} ${ns}:${service}
