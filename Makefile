GIT_COMMIT:=$(shell git rev-list -1 HEAD)
GH_USER:=krol3
GOOS=darwin
GOARCH=arm64

.PHONY: build

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X main.GitCommit=$(GIT_COMMIT)" -o bin/$(GOOS)_$(GOARCH)/$(GH_USER)
	echo "Build complete: ${GIT_COMMIT}"