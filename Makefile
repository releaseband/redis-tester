APP-NAME ?= redis-tester
GO-VERSION ?= 1.15.6

run:
	go run main.go

docker-run:
	docker run  \
		--rm -t -i \
		-v $(PWD):/app -w /app \
    	--name $(APP-NAME) \
		-v $(APP-NAME)-go-vol:/go \
		golang:$(GO-VERSION) \
		make run