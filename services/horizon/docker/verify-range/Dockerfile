FROM ubuntu:22.04

ARG GRAMR_VERSION 
ENV GRAMR_VERSION=${GRAMR_VERSION:-*}
# to remove tzdata interactive flow
ENV DEBIAN_FRONTEND=noninteractive

ADD dependencies /
RUN ["chmod", "+x", "dependencies"]
RUN /dependencies

ADD gramr.cfg /
ADD captive-core-pubnet.cfg /
RUN mkdir -p /cc
RUN mkdir -p /data

ADD start /
RUN ["chmod", "+x", "start"]

ENTRYPOINT ["/start"]
