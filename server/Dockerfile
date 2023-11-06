# build
FROM golang:1.21-alpine AS build-stage
WORKDIR /opt
COPY . .
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -o ./pkg/docs

RUN CGO_ENABLED=0 GOOS=linux go build -o /binary

# release
FROM alpine:3.18.4 AS release-stage
WORKDIR /opt
COPY --from=build-stage /binary /op/binary
CMD ["/opt/binary"]