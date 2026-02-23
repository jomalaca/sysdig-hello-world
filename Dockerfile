FROM golang:1.19 AS builder
#FROM golang:1.25.5-alpine AS builder


RUN mkdir /build
WORKDIR /build

COPY . .

RUN go get
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build hello-world.go

FROM scratch

LABEL maintainer="Josh Ma"
LABEL org.opencontainers.image.authors="Josh Ma <joshua.ma@sysdig.com>"
LABEL status="Testing"
LABEL description="This toy container is used for testing pipleline scanning in Sysdig Secure."

COPY --from=builder /build/hello-world /hello-world

ENTRYPOINT ["/hello-world"]
