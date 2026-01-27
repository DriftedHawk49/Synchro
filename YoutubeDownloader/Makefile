APP_NAME := YoutubeDownloader
IMAGE_NAME := ytdl
IMAGE_TAG := 1.0.0

GOOS := linux
GOARCH := amd64

.PHONY: all build docker-build docker-run clean

all: build

build:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) \
	go build -o $(APP_NAME) ./cmd/main.go

container: build
	sudo docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

clean:
	rm -f $(APP_NAME)
