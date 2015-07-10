FROM golang
 
ADD ~/var/apps/remindbot/* /go/src/github.com/aranair/remindbot/

RUN go install /go/src/github.com/aranair/remindbot

ADD ~/config.toml /go/bin/

ENTRYPOINT /go/bin/remindbot
 
EXPOSE 8080
