.PYHON: build
build:
	cd $(CURDIR)"/conf/mysql_1_master" && docker build -t mysql/master1 .
	cd $(CURDIR)"/conf/mysql_2_master" && docker build -t mysql/master2 .
	cd $(CURDIR)"/conf/mysql_1_slave1" && docker build -t mysql/slave1 .
	cd $(CURDIR)"/conf/mysql_2_slave1" && docker build -t mysql/slave2 .

.PYHON: deploy
deploy:
	docker stack deploy -c stack.yaml systemdesign

.PYHON: clean
clean:
	docker stack rm systemdesign