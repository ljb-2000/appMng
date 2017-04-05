FROM registry.time-track.cn:8052/luocheng/baseimg:1.0

MAINTAINER cheng.luo@hnair.com

COPY . /usr/bin/base

RUN chmod +x /usr/bin/base/go.sh
RUN /bin/sh /usr/bin/base/go.sh

ENV GOPATH=/root
ENV PATH=$PATH:/usr/local/go/bin

