FROM golang
 
ADD /var/apps/remindbot /go/src/github.com/aranair/remindbot
RUN go install github.com/aranair/remindbot

WORKDIR /go/src/github.com/aranair/remindbot
RUN go build -o build/remindbot ./app

ADD ~/config.toml /go/src/github.com/aranair/remindbot/build

ENTRYPOINT build/remindbot
 
EXPOSE 8080
