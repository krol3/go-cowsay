FROM golang:1.18-alpine AS builder
workdir /app
RUN --mount=type=bind,target=. \
    go build -ldflags "-s -w" -o /usr/bin/app ./server/server.go

FROM alpine:3.14
COPY --from=builder /usr/bin/app /usr/bin/app
CMD ["app"]