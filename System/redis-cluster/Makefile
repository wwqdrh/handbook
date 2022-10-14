SHELL := /bin/bash -o pipefail

.PHONY: build
build:
	cd $(CURDIR)"/proxy" && docker build --build-arg https_proxy=http://192.168.0.101:7890 -t wwqdrh/redis-proxy .

.PHONY: deploy
deploy:
	docker stack deploy -c stack.yaml redis-cluster
