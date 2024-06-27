# build
FROM golang:1.22.3 AS build-stage
WORKDIR /app
COPY . .
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -o ./pkg/docs

RUN CGO_ENABLED=0 GOOS=linux go build -o /binary

# release
FROM surnet/alpine-wkhtmltopdf:3.19.0-0.12.6-small AS release-stage
# FROM alpine:3.18.4 AS release-stage
RUN apk add --no-cache tini
WORKDIR /app
COPY --from=build-stage /binary /app/binary
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/app/binary"]