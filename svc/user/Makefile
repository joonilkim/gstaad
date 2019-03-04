root := ../..
ns := gstaad
service := user

test:
	go test

docker-build:
	docker build \
		-f Dockerfile \
		-t ${ns}:${service} ${root}

dev: docker-build
	docker run \
		-p 9200:9000 \
		--env-file .env --rm \
		--name ${ns}-${service} ${ns}:${service}

