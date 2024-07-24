FROM golang:1.22 AS builder
WORKDIR /src
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app-platform ./cmd/app-platform/

FROM alpine:latest AS production
WORKDIR /bin
COPY --from=builder /src/app-platform .

EXPOSE 8080

CMD ["app-platform"]
