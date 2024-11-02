GO_VERSION := go1.22

.PHONY: setup start start-dev tidy fmt tag test cover report coverage build curl start-prod

setup:
	@echo $(GO_VERSION)
	gvm install $(GO_VERSION)

start:
	go run cmd/main.go

tidy:
	go mod tidy

fmt:
	go fmt ./...

tag:
	git tag v0.0.1
	git push origin --tag

test:
	go test ./... --v

cover:
	go test ./... -coverprofile=coverage.out

report: cover
	go tool cover -html=coverage.out -o coverage.html

coverage:
	# Check if coverage meets 80% threshold
	@go tool cover -func coverage.out | grep "total:" | /usr/bin/awk '{print ((int($$3) > 80) != 1) }'

build:
	go build -o bin/main cmd/main.go

curl:
	curl http://localhost:8080/hello

start-prod: build
	./bin/main
