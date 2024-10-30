GO_VERSION := 1.22


setup:
	gvm install go$(GO_VERSION)
	gvm use go$(GO_VERSION)

build:
	go build -o api cmd/main.go

run:
	go run cmd/main.go