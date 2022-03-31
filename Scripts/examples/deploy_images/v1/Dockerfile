FROM golang:1.17-buster AS builder
ENV CGO_ENABLED 0
ENV GO111MODULE on
ENV GOFLAGS -mod=vendor
COPY . /go/src
WORKDIR /go/src
RUN go mod vendor
RUN go build -v -o echo .

# ===========================

FROM ubuntu
COPY --from=builder /go/src/echo /

WORKDIR /
ENV PARAMS=""
CMD ["sh","-c","/echo $PARAMS  > /dev/null 2>&1"]
