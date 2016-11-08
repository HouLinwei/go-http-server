FROM python:2.7.12-slim

MAINTAINER linwei.hou <linwei.hou@baifendian.com>

LABEL desc="build my first django environment"

RUN apt-get update && apt-get install -y --no-install-recommends \
			fortune

ENV DJANGO_VERION=1.10

ADD . /opt/

WORKDIR /opt

RUN pip install django==DJANGO_VERION

CMD ["python"]