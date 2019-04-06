FROM golang

RUN mkdir -p /go/src/github.com/angadsharma1016/grofers-task

ADD . /go/src/github.com/angadsharma1016/grofers-task

WORKDIR /go/src/github.com/angadsharma1016/grofers-task

RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

EXPOSE 3000

ENTRYPOINT watcher -run github.com/angadsharma1016/grofers-task/ -watch github.com/angadsharma1016/grofers-task