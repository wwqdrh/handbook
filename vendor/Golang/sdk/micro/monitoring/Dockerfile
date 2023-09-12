FROM golang:alpine

ENV GOPATH /code/
ADD . /code/src/wwqdrh/handbook/tools/micro0/monitoring
WORKDIR /code/src/wwqdrh/handbook/tools/micro0/monitoring
RUN go build

ENTRYPOINT /code/src/wwqdrh/handbook/tools/micro0/monitoring/monitoring
