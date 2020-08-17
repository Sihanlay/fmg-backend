#!/bin/bash
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
docker build -t registry.cn-hangzhou.aliyuncs.com/lsh-nc/lsh-nc-test:fmg-v1 .
docker push registry.cn-hangzhou.aliyuncs.com/lsh-nc/lsh-nc-test:fmg-v1