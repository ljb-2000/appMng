FROM alpine:3.4

MAINTAINER cheng.luo@hnair.com

ENV GOGS_URL=http://223.202.32.60:8071/
ENV UA_URL=

COPY appMng /usr/bin/appMng

EXPOSE 8088

ENTRYPOINT ["/usr/bin/appMng"]



