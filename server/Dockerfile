# build
FROM golang:1.21-alpine AS build-stage
WORKDIR /app
COPY . .
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -o ./pkg/docs

RUN CGO_ENABLED=0 GOOS=linux go build -o /binary

# release
FROM alpine:3.18.4 AS release-stage
RUN apk add --no-cache tini
WORKDIR /app
COPY --from=build-stage /binary /app/binary
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/app/binary"]