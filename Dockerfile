#FROM golang:1.18-alpine as builder
FROM golang:1.18.3-alpine as builder

LABEL maintainer="Sysdig CSE - dustin.krysak+maintainer@sysdig.com"
LABEL org.opencontainers.image.authors="dustin.krysak+maintainer@sysdig.com"


RUN mkdir /build
WORKDIR /build

COPY . .

RUN go get
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build hello-world.go

FROM scratch

COPY --from=builder /build/hello-world /hello-world

ENTRYPOINT ["/hello-world"]