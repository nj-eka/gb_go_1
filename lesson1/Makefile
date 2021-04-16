# https://habr.com/ru/post/461467/

PROJECT?=github.com/nj-eka/gb_go_1/lesson1

GOOS?=linux
GOARCH?=amd64

RELEASE := $(shell git tag -l | tail -1 | grep -E "v.+"|| echo devel)
COMMIT := git-$(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')

build:
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 go build \
		-ldflags "-s -w \
		-X ${PROJECT}/version.Version=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o bin/gb_go_1/lesson1 ${PROJECT}