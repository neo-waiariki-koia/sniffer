.PHONY: build image build-image run stop clean

build: $(shell find . -iname '*.go')
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o bin/sniffer-linux main.go

image: Dockerfile bin/sniffer-linux
	docker image build -t sniffer:${VERSION} .

build-image: build image

run:
	docker run -d --name sniffer --net haproxy -p 9090:9090 sniffer:${VERSION}

stop:
	docker stop sniffer; docker rm sniffer

clean:
	rm -rf bin
