
APP_ROOT ?= $(shell 'pwd')
IMAGE_NAME ?= robusdockerhub/pkg-base-image
# Alias command for docker's `make` executable
DOCKER_RUN ?=  \
	       docker run \
	       --rm \
	       -v $(APP_ROOT):/app \
	       -w /app \
	       $(IMAGE_NAME)
schema:
	go run github.com/99designs/gqlgen generate

docker/stop:
	docker-compose down

docker/start:
	docker-compose up --build --remove-orphans

docker/stop-start:
	make stop-local-docker && make start-local-docker

## Run pre-commit
docker/pre-commit:
	@$(DOCKER_RUN) pre-commit-run

## Install pre-commit
docker/pre-commit-install:
	@$(DOCKER_RUN) pre-commit-install

## Uninstall pre-commit
docker/pre-commit-uninstall:
	@$(DOCKER_RUN) pre-commit-uninstall

## Run golangci-lint
docker/golangci-lint:
	@$(DOCKER_RUN) golangci-lint-run

pre-commit-install:
	pre-commit install

pre-commit-uninstall:
	pre-commit uninstall

pre-commit-run:
	pre-commit run

golangci-lint-run:
	golangci-lint run
