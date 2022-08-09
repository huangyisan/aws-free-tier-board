.PHONY: build-linux build-osx build-windows clean

PWD=$(shell pwd)

VERSION=$(shell git rev-parse --short HEAD)
BUILD=$(shell date +%FT%T%z)

build-linux:
	@GOARCH=amd64 CGO_ENABLED=1 GOOS=linux go build -ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}" -o trojanTrafficCollector

build-osx:
	@GOARCH=amd64 CGO_ENABLED=1 GOOS=darwin go build -ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}" -o trojanTrafficCollector

build-windows:
	@GOARCH=amd64 CGO_ENABLED=1 GOOS=windows go build -ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}" -o trojanTrafficCollector

clean:
	@if [ -f trojanTrafficCollector ] ; then rm trojanTrafficCollector ; fi
