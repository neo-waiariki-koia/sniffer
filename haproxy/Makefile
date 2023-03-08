.PHONY: run stop

run:
	docker run -d --name haproxy --net haproxy -v $(shell pwd):/usr/local/etc/haproxy:ro -p 80:80 -p 8404:8404 haproxytech/haproxy-alpine:2.7

stop:
	docker stop haproxy; docker rm haproxy
