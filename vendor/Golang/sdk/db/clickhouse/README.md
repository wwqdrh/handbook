列式数据库

```shell
docker pull yandex/clickhouse-server:22.1.3.7-alpine

docker run -d --name clickhouse-server --ulimit nofile=262144:262144 \
-p 8123:8123 -p 9000:9000 -p 9009:9009 yandex/clickhouse-server:22.1.3.7-alpine
```

## 引擎

重新启动服务器时，表中的数据消失，表将变为空。通常，使用此表引擎是不合理的。但是，它可用于测试，以及在相对较少的行（最多约100,000,000）上需要最高性能的查询。