FROM alpine:latest AS build

RUN apk add --no-cache gcc linux-headers libc-dev make git

WORKDIR /tmp

RUN git clone --depth 1 --branch 1.0-beta2 https://github.com/RedisLabs/redis-cluster-proxy.git

RUN cd /tmp/redis-cluster-proxy && make && make install

RUN rm -Rf /tmp/redis-cluster-proxy

FROM alpine:latest

COPY --from=build /usr/local/bin/redis-cluster-proxy /usr/local/bin/redis-cluster-proxy