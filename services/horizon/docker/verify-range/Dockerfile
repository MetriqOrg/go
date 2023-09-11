FROM ubuntu:22.04

ARG GRAVITY_VERSION 
ENV GRAVITY_VERSION=${GRAVITY_VERSION:-*}
# to remove tzdata interactive flow
ENV DEBIAN_FRONTEND=noninteractive

ADD dependencies /
RUN ["chmod", "+x", "dependencies"]
RUN /dependencies

ADD gravity.cfg /
ADD captive-core-pubnet.cfg /
RUN mkdir -p /cc
RUN mkdir -p /data

ADD start /
RUN ["chmod", "+x", "start"]

ENTRYPOINT ["/start"]
