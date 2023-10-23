# The base image, overridable with arguments
ARG BASE_IMAGE=golang:1.21

# Install docker from docker image
FROM docker:24-cli AS docker

# The CI image base
FROM $BASE_IMAGE

LABEL description="Image for GitHub CI base workflow"

# Go binaries install path
ENV GOBIN=/usr/local/bin
WORKDIR /usr/local/bin

# Install required packages
RUN apt -qq update && \
    apt -qqy install \
        curl \
        wget \
        cython3 \
        pip \
        python3 \
        python3-dev \
        python3-pip \
        python3-venv \
        python3-setuptools \
        python3-grpcio \
        python3-grpc-tools \
        python3-protobuf \
        build-essential \
        zlib1g-dev \
        libncurses5-dev \
        libgdbm-dev \
        libnss3-dev \
        libssl-dev \
        libsqlite3-dev \
        libreadline-dev \
        libffi-dev \
        libluajit-5.1-dev \
        libbz2-dev && \
    rm -rf /root/.cache && apt-get -y autoremove && apt-get clean

# Install docker and docker compose tooling from docker image
COPY --from=docker /usr/local/libexec/docker/cli-plugins/docker-compose /usr/local/libexec/docker/cli-plugins/
COPY --from=docker /usr/local/bin/docker /usr/local/bin/
COPY --from=docker /usr/local/bin/docker-compose /usr/local/bin/

RUN docker compose version
RUN go version
RUN python3 -V
