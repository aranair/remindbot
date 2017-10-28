FROM golang:1.9

ADD configs.toml /go/bin/

ADD . /go/src/github.com/aranair/remindbot
WORKDIR /go/src/github.com/aranair/remindbot

# RUN go get ./...
RUN go get github.com/tools/godep
RUN godep restore
RUN go install ./...

WORKDIR /go/src/github.com/aranair/remindbot
WORKDIR /go/bin/
