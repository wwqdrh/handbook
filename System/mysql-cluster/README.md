# 简介

基于mysql8的内置集群功能，mysql8还有sh工具简化操作

multi-master模式，一个主节点坏了自动选择一个新的主节点

搭配mysql-router做路由代理

# 环境搭建

> 某个容器停止能够自动重新选组，但是如果机器重启集群会坏

## 基础配置

基础环境准备

`docker pull mysql/mysql-server:8.0`

network创建

`docker network create --driver overlay mysqlcluster`

volume创建

多个mysql使用的持久目录

`docker volume create mysqlcluster`

启动容器

```bash
for N in 1 2 3 4
  do docker service create --name=mysql$N --hostname=mysql$N --network=mysqlcluster --mount source=myvol2,target=/var/lib/mysql \
      -e MYSQL_ROOT_PASSWORD=root mysql/mysql-server:8.0
done
```

配置用户权限

```bash
for N in 1 2 3 4
do docker exec -it `docker ps | grep mysql$N | awk '{print $1}'` mysql -uroot -proot \
  -e "CREATE USER 'inno'@'%' IDENTIFIED BY 'inno';" \
  -e "GRANT ALL privileges ON *.* TO 'inno'@'%' with grant option;" \
  -e "reset master;"
done
```

校验用户是否创建成功

```bash
for N in 1 2 3 4
do docker exec -it `docker ps | grep mysql$N | awk '{print $1}'` mysql -uinno -pinno \
  -e "SHOW VARIABLES WHERE Variable_name = 'hostname';" \
  -e "SELECT user FROM mysql.user where user = 'inno';"
done
```

## 配置节点信息

`docker exec -it [id] mysqlsh -uroot -proot -S/var/run/mysqld/mysqlx.sock`

检查 `dba.checkInstanceConfiguration("inno@mysql1:3306")`

设置集群 `dba.configureInstance("inno@mysql1:3306")`

and do mysql2 3 4

重启服务(必须重启，所以这个容器记得持久化)

## 创建集群

`docker exec -it mysql1 mysqlsh -uroot -proot -S/var/run/mysqld/mysqlx.sock`

`var cluster = dba.createCluster("mycluster")`

`cluster.status()`

`cluster.describe()`

添加其他节点

`cluster.addInstance("inno@mysql2:3306")`

## 启动router组件

```bash
docker run -d --name mysql-router --net=innodbnet \
   -e MYSQL_HOST=mysql1 \
   -e MYSQL_PORT=3306 \
   -e MYSQL_USER=inno \
   -e MYSQL_PASSWORD=inno \
   -e MYSQL_INNODB_CLUSTER_MEMBERS=4 \
   mysql/mysql-router
```

done!
