FROM ubuntu:focal

ARG VERSION
ARG GRAVITY_VERSION
ARG DEBIAN_FRONTEND=noninteractive
ARG ALLOW_CORE_UNSTABLE=no

RUN apt-get update && apt-get install -y wget apt-transport-https gnupg2 && \
    wget -qO /etc/apt/trusted.gpg.d/SDF.asc https://apt.stellar.org/SDF.asc && \
    echo "deb https://apt.stellar.org focal stable" | tee -a /etc/apt/sources.list.d/SDF.list && \
    if [ "${ALLOW_CORE_UNSTABLE}" = "yes" ]; then echo "deb https://apt.stellar.org focal unstable" | tee -a /etc/apt/sources.list.d/SDF.list; fi && \
    cat /etc/apt/sources.list.d/SDF.list && \
    apt-get update && apt-cache madison gravity && eval "apt-get install -y gravity${GRAVITY_VERSION+=$GRAVITY_VERSION}" && \
    if [ "${ALLOW_CORE_UNSTABLE}" = "yes" ]; then sed -i '/unstable/d' /etc/apt/sources.list.d/SDF.list; fi && \
    cat /etc/apt/sources.list.d/SDF.list && \
    echo "deb https://apt.stellar.org focal testing" | tee -a /etc/apt/sources.list.d/SDF.list && \
    apt-get update && apt-cache madison stellar-orbitr && apt-get install -y stellar-orbitr=${VERSION} && \
    apt-get clean && rm -rf /var/lib/apt/lists/* /var/log/*.log /var/log/*/*.log

EXPOSE 8000
ENTRYPOINT ["/usr/bin/stellar-orbitr"]
