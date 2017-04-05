#!/bin/bash

env GOOS=linux go build
cp appMng pkg
cd pkg
docker build -t registry.time-track.cn:8052/luocheng/appmng:0.9 .
docker push registry.time-track.cn:8052/luocheng/appmng:0.9