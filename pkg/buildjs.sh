#!/bin/sh
git clone {{.Git}}
cp -a template/JavaScript/. {{.App}}
docker build -t {{.Reg}}/{{.User}}/{{.Img}}:{{.Tag}} ./{{.App}}
docker login -u innovation -p "0p;/(OL>" https://{{.Reg}}
docker push {{.Reg}}/{{.User}}/{{.Img}}:{{.Tag}}