.PHONY: compile build push deploy run clean

IMAGE = go-hello-world
VERSION = latest

all: compile build run

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(IMAGE) main.go

build:
	docker build -t $(IMAGE):$(VERSION) .

run:
	docker run -p 3000:3000 $(IMAGE)

clean:
	docker system prune
