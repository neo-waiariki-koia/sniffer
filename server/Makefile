.PHONY: build image build-image run stop clean

build: $(shell find . -iname '*.go')
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o bin/simple-webserver-linux main.go

image: Dockerfile bin/simple-webserver-linux
	docker image build -t server:${VERSION} .

build-image: build image

run:
	docker run -d --name server --net haproxy -p 8080:8080 server:${VERSION}

stop:
	docker stop server; docker rm server

clean:
	rm -rf bin
