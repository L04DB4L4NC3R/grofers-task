FROM golang

RUN mkdir -p /go/src/github.com/angadsharma1016/grofers-task

ADD . /go/src/github.com/angadsharma1016/grofers-task
WORKDIR /go/src/github.com/angadsharma1016/grofers-task/server

EXPOSE 3000
ENTRYPOINT go run main.go