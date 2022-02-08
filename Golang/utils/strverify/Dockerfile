# 编译：docker build --no-cache -t tpler .
# 运行：docker run --rm -it -p 8080:8080 tpler
# 导出：docker save tpler > tpler.tar
# 导入：docker load < tpler.tar

FROM golang:1.17 AS strverify
ENV GO111MODULE on
ENV GOFLAGS -mod=vendor
COPY . /go/src
WORKDIR /go/src
RUN go build -v -o strverify -ldflags "-s -w -linkmode 'external' -extldflags '-static'" ./cli/main.go

# ===========================

FROM centos:centos8
COPY --from=strverify /go/src/strverify /

WORKDIR /
ENV PRODUCTION=1
ENV PARAMS=""
RUN ["chmod", "+x", "/strverify"]
CMD ["sh","-c","/strverify $PARAMS"]