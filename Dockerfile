FROM golang:1.6

ADD configs.toml /go/bin/

ADD . /go/src/github.com/aranair/remindbot
WORKDIR /go/src/github.com/aranair/remindbot

RUN go get ./...
RUN go install ./...

WORKDIR /go/bin/
ENTRYPOINT remindbot

EXPOSE 8080
