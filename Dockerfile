#FROM golang:1.18-alpine as builder
FROM golang:1.25.5-alpine AS builder


RUN mkdir /build
WORKDIR /build

COPY . .

RUN go get
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build hello-world.go

FROM scratch

LABEL maintainer="Dustin krysak"
LABEL org.opencontainers.image.authors="Dustin Krysak <dustin.krysak+maintainer@sysdig.com>"
LABEL status="testing"
LABEL description="This toy container is used for testing pipleline scanning in Sysdig Secure."

COPY --from=builder /build/hello-world /hello-world

ENTRYPOINT ["/hello-world"]
