SHELL := /bin/bash -o pipefail

.PHONY: build
build:
	cd $(CURDIR)"/sential" && docker build -t redis:6-sential .

.PHONY: deploy
deploy:
	docker stack deploy -c stack.yaml redis-sential
