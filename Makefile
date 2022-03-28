ROOT_APP_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
VERSION=`git describe --tags --dirty --always`
COMMIT=`git rev-parse HEAD`
LDFLAGS=-ldflags "-w -s -X main.version=${VERSION} -X main.commit=${COMMIT}"

.PHONY: run
run:
	docker-compose up --build && cd ./cmd/anyData && go run main.go
