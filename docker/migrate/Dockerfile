FROM debian:buster

WORKDIR /migrate

RUN apt-get update && apt-get install -y \
    curl \
    build-essential \
&& apt-get clean \
&& rm -rf /var/lib/apt/lists/*

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.13.0/migrate.linux-amd64.tar.gz | tar xvz
RUN mv ./migrate.linux-amd64 /usr/bin/migrate

RUN mkdir -p ./migrations/
COPY ./db/migrations ./migrations/
COPY ./docker/wait-for-it.sh /usr/bin/wait-for-it
COPY ./help.mk ./
COPY ./docker/migrate/Makefile ./
