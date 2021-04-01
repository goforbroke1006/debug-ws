build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o debug-ws .

test:
	go test -v ./...

lint:
	echo 'pls setup linter'

DOCKER_REGISTRY=docker.io
DOCKER_OWNER=goforbroke1006
DOCKER_IMAGE_NAME=debug-ws

image:
	docker build -t ${DOCKER_REGISTRY}/${DOCKER_OWNER}/${DOCKER_IMAGE_NAME}:latest .
	docker push ${DOCKER_REGISTRY}/${DOCKER_OWNER}/${DOCKER_IMAGE_NAME}:latest