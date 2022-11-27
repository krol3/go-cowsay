GIT_COMMIT:=$(shell git rev-list -1 HEAD)
GH_APP:=go-cowsay
GOOS=darwin
GOARCH=arm64

.PHONY: build

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X main.GitCommit=$(GIT_COMMIT)" -o bin/$(GOOS)_$(GOARCH)/$(GH_APP) ./server/server.go
	echo "Build complete: ${GIT_COMMIT}"

test:
	go test -v ./...

image:
	docker buildx build -t $(GH_APP):$(GIT_COMMIT) .