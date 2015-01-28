# Image: justrunit/php

FROM justrunit/debian

MAINTAINER Amarnath Ravikumar <amar@gmail.com>

USER root

RUN apt-get update && \
	apt-get install -y php5-cli && \
	apt-get clean && \
	rm -rf /var/lib/apt/lists/* && \
	echo "date.timezone = \"UTC\"" >> /etc/php5/cli/php.ini && \
	echo "memory_limit = 128M" >> /etc/php5/cli/php.ini && \
	curl -sS https://getcomposer.org/installer | php && \
	mv composer.phar /usr/local/bin/composer && \
	mkdir /home/justrunit/services/myproject && \
	chown justrunit:justrunit -R /home/justrunit/services/myproject

VOLUME ["/home/justrunit/services/myproject"]

USER justrunit

WORKDIR /home/justrunit/services/myproject

CMD if [ -s "composer.json" ]; then composer install --verbose; fi; \
	php code

ENV LINT_CMD php -l
