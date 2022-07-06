#FROM golang:1.18-alpine as builder
FROM golang:1.18.3-alpine as builder

RUN mkdir /build
WORKDIR /build

COPY . .

RUN go get
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build hello-world.go
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o hello-world

FROM scratch

COPY --from=builder /build/hello-world /hello-world

ENTRYPOINT ["/hello-world"]