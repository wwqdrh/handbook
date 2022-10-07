# 实验环境搭建

mysql、redis这些集群都是ap模型

主库配置

```bash
GRANT REPLICATION SLAVE ON *.* to 'replicer'@'%' identified by '123456';

# or
create user 'replicer'@'%' identified by '123456';  
grant replication slave on *.* to 'replicer'@'%';  
flush privileges;

# 查看主容器数据库状态，记录File的值和Position的值
show master status;
```

从库配置

```bash
change master to master_host='mysql_2_master', master_port=3306, master_user='replicer', master_password='123456', master_log_file='mysql-bin.000003', master_log_pos=441;

start slave;

# 检查主从连接状态
show slave status;

show slave status\G; # 更好看


# 终止slave
STOP SLAVE IO_THREAD FOR CHANNEL '';

# 如果出错
set global sql_slave_skip_counter=1;

start slave;

# 将MySQL设置为只读状态的命令
mysql> show global variables like "%read_only%";
mysql> flush tables with read lock;
mysql> set global read_only=1;
mysql> show global variables like "%read_only%";

# 将MySQL从只读设置为读写状态的命令
mysql> unlock tables;
mysql> set global read_only=0;
```

# 代码测试

主库写数据

从库读数据

如果产生分区，那么虽然可用，但是数据一致性会有问题，需要引入额外方法来减少这个数据不一致带来的影响(如果是影响不大则可忽略)