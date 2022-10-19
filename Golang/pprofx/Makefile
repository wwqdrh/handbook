SHELL := /bin/bash pipefail

.PHONY: env
env:
	docker pull pyroscope/pyroscope:latest

	docker service create --name pyrocsope --network dev pyroscope/pyroscope:latest server