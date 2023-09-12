go语言中使用sqlite3必须开启CGO才能运行

并且需要保证环境下包含sqlite环境才能使用（比python麻烦太多了）

## 示例

``` bash
FROM golang:alpine AS build

# Important:
#   Because this is a CGO enabled package, you are required to set it as 1.
ENV CGO_ENABLED=1

RUN apk add --no-cache \
    # Important: required for go-sqlite3
    gcc \
    # Required for Alpine
    musl-dev

WORKDIR /workspace

COPY . /workspace/

RUN \
    go mod init github.com/mattn/sample && \
    go mod tidy && \
    go install -ldflags='-s -w -extldflags "-static"' ./simple.go

RUN \
    # Smoke test
    set -o pipefail; \
    /go/bin/simple | grep 99\ こんにちわ世界099

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM scratch

COPY --from=build /go/bin/simple /usr/local/bin/simple

ENTRYPOINT [ "/usr/local/bin/simple" ]
```