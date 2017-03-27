FROM alpine:3.4

MAINTAINER cheng.luo@hnair.com

ENV GOGS_URL=http://223.202.32.60:8071/
ENV UA_URL=http://172.16.5.240:8072/
ENV K8S_URL=http://172.16.5.245:8080/

COPY appMng /usr/bin/appMng

EXPOSE 8088

ENTRYPOINT ["/usr/bin/appMng"]



