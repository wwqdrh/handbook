#!/bin/sh

C_DIR=$(
    cd $(dirname $0)
    pwd
)

C_V1=$C_DIR"/v1"

C_V2=$C_DIR"/latest"

function main() {
    deploy_v1
    deploy_v2
}

function deploy_v1() {
    cd $C_V1
    docker build . -t 192.168.110.114:5000/library/deployecho:v1
    docker push 192.168.110.114:5000/library/deployecho:v1
}

function deploy_v2() {
    cd $C_V2
    docker build . -t 192.168.110.114:5000/library/deployecho:latest
    docker push 192.168.110.114:5000/library/deployecho:latest
}

main $@
