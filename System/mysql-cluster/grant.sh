#!/bin/bash

for N in 1 2 3 4
do docker exec -it `docker ps | grep mysql_mysql$N | awk '{print $1}'` mysql -uroot -proot \
  -e "CREATE USER 'inno'@'%' IDENTIFIED BY 'inno';" \
  -e "GRANT ALL privileges ON *.* TO 'inno'@'%' with grant option;" \
  -e "reset master;"
done

for N in 1 2 3 4
do docker exec -it `docker ps | grep mysql_mysql$N | awk '{print $1}'` mysql -uinno -pinno \
  -e "SHOW VARIABLES WHERE Variable_name = 'hostname';" \
  -e "SELECT user FROM mysql.user where user = 'inno';"
done