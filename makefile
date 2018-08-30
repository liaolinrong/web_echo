all: push

TAG = 0.1
PREFIX = registry.hundsun.com/hcs/web_echo

server: web_echo.go
	CGO_ENABLED=0 go build -a -installsuffix cgo --ldflags '-w' ./web_echo.go

container: server
#	docker build -f Dockerfile-scratch -t $(PREFIX):$(TAG) .
	docker build -t $(PREFIX):$(TAG) .

push: container
	docker push $(PREFIX):$(TAG)

clean:
	rm -f web_echo
