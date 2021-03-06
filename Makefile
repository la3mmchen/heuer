#!make
include properties.env
export $(shell sed 's/=.*//' properties.env)
GIT_COMMIT := $(shell git describe --always --long --dirty)
PROJECT_NAME := $(shell basename "$$PWD")

.DEFAULT_GOAL := default

test:
	echo "${PROJECT_NAME}"

.PHONY: default
default: build run-help run-list

.PHONY: build
build: 
	@rm -f ${EXECUTABLE}
	@go build -o ${EXECUTABLE} -ldflags "-X main.AppVersion=${GIT_COMMIT}" .

.PHONY: docker
docker: 
	docker build --build-arg app_version=$(GIT_COMMIT) -t ${PROJECT}/${PROJECT_NAME}:${GIT_COMMIT} .
	docker tag ${PROJECT}/${PROJECT_NAME}:${GIT_COMMIT} ${PROJECT}/${PROJECT_NAME}:latest

run-help:
	@./${EXECUTABLE} --help

run-list:
	@./${EXECUTABLE} read

docker-run:
	docker run --rm -it ${PROJECT}/${PROJECT_NAME}:latest --help

prom: 
	@echo "\n"
	docker run -t --rm -v $(PWD)/data:/data --entrypoint=/bin/sh --workdir=/data prom/prometheus:latest -c "/bin/promtool test rules test-example-rules.yaml"