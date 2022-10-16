#build stage
FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN make build

#final stage
FROM alpine:3.10
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/bin /app
RUN mv /app/go-cowsay* /app/cowsay
RUN chmod +x /app/cowsay
ENTRYPOINT /app/cowsay
EXPOSE 8005