```bash
docker service create --name pyroscope -p 4040:4040 --network dev --mount type=volume,source=pyroscope_config,dst=/etc/pyroscope pyroscope/pyroscope:latest server

# TODO 下面这个写入配置暂时没找到自动脚本
conf=$(cat <<EOF
---
log-level: debug
scrape-configs:
  - job-name: testing
    enabled-profiles: [cpu, mem, goroutines, mutex, block]
    static-configs:
      - application: test-app
        spy-name: gospy
        targets:
          - code-server:6060
        labels:
          env: dev
EOF
)

docker exec -it `docker ps | grep pyroscope | awk '{print $1}'` bash -c 'echo "$conf" >> /etc/pyroscope/server-simple.yml'

# 手动操作
docker exec -it `docker ps | grep pyroscope | awk '{print $1}'` bash

docker exec -it `docker ps | grep pyroscope | awk '{print $1}'` ls /etc/pyroscope/

docker service update --args 'server --config /etc/pyroscope/server.simple.yml' pyroscope
```

second time start can direct set yml

`docker service create --name pyroscope -p 4040:4040 --network dev --mount type=volume,source=pyroscope_config,dst=/etc/pyroscope pyroscope/pyroscope:latest server --config /etc/pyroscope/server.simple.yml`