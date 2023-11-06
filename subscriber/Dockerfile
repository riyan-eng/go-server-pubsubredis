# build
FROM golang:1.21-alpine AS build-stage
WORKDIR /opt
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /binary

# release
FROM alpine:3.18.4 AS release-stage
WORKDIR /opt
COPY --from=build-stage /binary /opt/binary
CMD ["/opt/binary"]