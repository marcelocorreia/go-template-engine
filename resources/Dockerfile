FROM alpine:3.9
MAINTAINER Marcelo Correia <marcelo@correia.io>

ARG gte_version="3.1.6"
WORKDIR /usr/local/bin/
RUN apk update
RUN apk add --no-cache --update \
	unzip
ADD ./dist/go-template-engine-linux-amd64-${gte_version}.zip /usr/local/bin/go-template-engine.zip
RUN unzip go-template-engine.zip
RUN chmod 0755 /usr/local/bin/go-template-engine

RUN apk del unzip
RUN rm /var/cache/apk/*

RUN mkdir /app
WORKDIR /app
ENTRYPOINT ["/usr/local/bin/go-template-engine"]
