GIT_COMMIT:=$(shell git rev-list -1 HEAD)
DOMAIN=krol
GH_APP:=go-cowsay
GOOS=darwin
GOARCH=arm64

.PHONY: build

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X main.GitCommit=$(GIT_COMMIT)" -o bin/$(GH_APP)_$(GOOS)_$(GOARCH) ./server/server.go
	echo "Build complete: ${GIT_COMMIT}"

test:
	go test -v ./...

build-image:
	docker build -t ${DOMAIN}/${GH_APP}:${GIT_COMMIT} .

push-image:
	docker push ${DOMAIN}/${GH_APP}:${GIT_COMMIT}

sign-image:
	cosign sign --key ./keys/cosign.key -a signed=krol ${DOMAIN}/${GH_APP}:${GIT_COMMIT}