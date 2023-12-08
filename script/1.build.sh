#!/bin/bash

docker_id="ketidevit2"
image_name="exascale.exascale-web"
operator="exascale.exascale-web"
version="latest"

export GO111MODULE=on
go mod vendor

go build -o ../build/_output/bin/$operator -mod=vendor ../main.go
docker build -t $docker_id/$image_name:$version ../build
docker push $docker_id/$image_name:$version
