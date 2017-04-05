FROM ubuntu:14.04

MAINTAINER cheng.luo@hnair.com

COPY . /usr/bin/base

RUN chmod +x /usr/bin/base/docker.sh
RUN /bin/sh /usr/bin/base/docker.sh
RUN apt-get install git



