FROM golang:1.4

RUN go get github.com/aranair/remindbot
WORKDIR /go/src/github.com/aranair/remindbot
RUN go get ./...
RUN go install ./...

ADD configs.toml /go/bin/

ENTRYPOINT /go/bin/remindbot

EXPOSE 8080
