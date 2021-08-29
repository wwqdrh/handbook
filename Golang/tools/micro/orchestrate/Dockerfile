FROM golang:alpine

ENV GOPATH /code/
ADD . /code/src/wwqdrh/handbook/tools/micro0/docker
WORKDIR /code/src/wwqdrh/handbook/tools/micro0/docker/example
RUN go build

ENTRYPOINT /code/src/wwqdrh/handbook/tools/micro0/docker/example/example


