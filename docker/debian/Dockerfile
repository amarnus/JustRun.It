# Image: justrunit/debian

FROM debian:wheezy

MAINTAINER Srinivasan Kidambi <kidambisrinivas@gmail.com>

ENV DEBIAN_FRONTEND noninteractive

ADD s6-1.1.3.2-musl-static.tar /
ADD huptime_0.1-29.g2d47374_amd64.deb /
ADD .ssh/id_rsa /
ADD .ssh/id_rsa.pub /

# 1. Install curl, git, huptime and S6
RUN export DEBIAN_FRONTEND=noninteractive && \
    echo 'deb http://ftp.us.debian.org/debian/ testing main contrib non-free' >> /etc/apt/sources.list && \
    apt-get update && \
    apt-get install -t testing libc6 && \
    apt-get install -y curl git python wget g++ gcc make && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    dpkg -i /huptime_0.1-29.g2d47374_amd64.deb && \
    mkdir -p /etc/s6/.s6-svscan && \
    ln -s /bin/true /etc/s6/.s6-svscan/finish

# 2. Setup the system

RUN ln -sf /bin/bash /bin/sh && \
    addgroup --gid 1001 justrunit && \
    useradd -u 1001 -d /home/justrunit -m justrunit -g justrunit && \
    echo 'justrunit:justrunit' | chpasswd && \
    adduser justrunit justrunit && \
    chsh -s /bin/bash justrunit && \
    mkdir -p /etc/ssh/ && \
    echo "    IdentityFile /home/justrunit/.ssh/id_rsa" >> /etc/ssh/ssh_config && \
    mkdir -p /home/justrunit/.ssh && \
    echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /home/justrunit/.ssh/config && \
    cp /id_rsa /home/justrunit/.ssh/id_rsa && \
    cp /id_rsa.pub /home/justrunit/.ssh/id_rsa.pub && \
    chown -R justrunit:justrunit /home/justrunit/.ssh

# 3. Setup directories for services 

USER justrunit
ENV HOME /home/justrunit

# Add the SSH keys to clone Justrunit repositories
RUN chmod 700 /home/justrunit/.ssh && \

    # Set permissions to the SSH keys
    chmod 600 /home/justrunit/.ssh/id_rsa && \
    chmod 600 /home/justrunit/.ssh/config && \
    chmod 644 /home/justrunit/.ssh/id_rsa.pub && \

    # Create volume directories for services
    mkdir /home/justrunit/services && \

    # Create volume directories for jobs
    mkdir /home/justrunit/jobs && \

    # Create volume directories for databases
    mkdir /home/justrunit/data && \

    # Create volume directories for logs
    mkdir /home/justrunit/logs

WORKDIR /home/justrunit

