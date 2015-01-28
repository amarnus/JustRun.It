# Image: justrunit/python

FROM justrunit/debian

MAINTAINER Srinivasan Kidambi <kidambisrinivas@gmail.com>

USER root

RUN set -x && \
	apt-get update && \
	apt-get install -y python \
	python-dev \
	python-pip \
	python-virtualenv && \
    apt-get clean && \
	rm -rf /var/lib/apt/lists/* && \
	easy_install pip && \
	pip install virtualenv && \
	pip install pyflakes && \
	mkdir /home/justrunit/services/myproject && \
	chown justrunit:justrunit -R /home/justrunit/services/myproject

VOLUME ["/home/justrunit/services/myproject"]

USER justrunit

WORKDIR /home/justrunit/services/myproject
ADD ./pydefaults /home/justrunit/pydefaults
RUN virtualenv env && mv env /home/justrunit

CMD rm -rf ./env && cp -r /home/justrunit/env ./ && \
	sort /home/justrunit/pydefaults > pyd && \
	mv -f pyd /home/justrunit/pydefaults && \
	sort requirements.txt > sr && \
	comm -13 /home/justrunit/pydefaults sr > reqt && \
	mv reqt requirements.txt && \
	if [ -s "requirements.txt" ]; then ./env/bin/pip install --timeout 5 -v -r requirements.txt; fi; \
	env/bin/python code

ENV LINT_CMD pyflakes
