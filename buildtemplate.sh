#!/bin/sh
git clone http://223.202.32.60:8071/{{.User}}/{{.App}}.git
cp -a template/go/. {{.App}}
env GOOS=linux go build
docker build -t registry.time-track.cn:8443/{{.User}}/{{.App}}:{{.Tag}} ./{{.App}}
docker login -u admin -p "0p;/(OL>" https://registry.time-track.cn:8443
docker push registry.time-track.cn:8443/{{.User}}/{{.App}}:{{.Tag}}