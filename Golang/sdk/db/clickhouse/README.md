列式数据库

```shell
docker pull yandex/clickhouse-server:22.1.3.7-alpine

docker run -d --name clickhouse-server --ulimit nofile=262144:262144 \
-p 8123:8123 -p 9000:9000 -p 9009:9009 yandex/clickhouse-server:22.1.3.7-alpine
```