# Image: justrunit/node

FROM justrunit/debian

MAINTAINER Amarnath Ravikumar <amarnus@gmail.com>

USER root

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && \
    curl -sL https://deb.nodesource.com/setup | bash -

RUN apt-get update && \
    apt-get install -y nodejs && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    npm install -g jshint && \
    mkdir /home/justrunit/services/myproject && \
    chown justrunit:justrunit -R /home/justrunit/services/myproject

RUN mkdir -p ~/.npm && chown -R justrunit ~/.npm

VOLUME ["/home/justrunit/services/myproject"]

USER justrunit

WORKDIR /home/justrunit/services/myproject

CMD if [ -s "deps" ]; then cat deps | xargs -l npm install; fi; \
	node code

ENV LINT_CMD jshint
