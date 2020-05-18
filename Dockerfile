FROM golang:1.14 as builder

RUN mkdir -p /go/src/github.com/hsmtkk/heroku-go-web-app

WORKDIR /go/src/github.com/hsmtkk/heroku-go-web-app

COPY ./go.mod .
COPY ./go.sum .
COPY ./pkg/helloworld ./pkg/helloworld
COPY ./cmd/helloworld ./cmd/helloworld

WORKDIR /go/src/github.com/hsmtkk/heroku-go-web-app/cmd/helloworld

ENV CGO_ENABLED=0

RUN go build -o /helloworld

FROM alpine:3.11.6

COPY --from=builder /helloworld /helloworld
COPY ./revision /revision

ENTRYPOINT ["/helloworld"]

