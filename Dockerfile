FROM golang:1.6

ADD configs.toml /go/bin/

ADD . /go/src/github.com/aranair/remindbot
WORKDIR /go/src/github.com/aranair/remindbot

# RUN go get ./...
RUN go get github.com/tools/godep
RUN godep restore
RUN go install ./...
RUN go get github.com/rubenv/sql-migrate/...

WORKDIR /go/src/github.com/aranair/remindbot
RUN sql-migrate up

WORKDIR /go/bin/
