FROM alpine:3.9
MAINTAINER Marcelo Correia <marcelo@correia.io>

WORKDIR /usr/local/bin/
RUN apk update
RUN apk add --no-cache --update \
	unzip
ADD ./dist/go-template-engine-linux-amd64-2.5.8.zip /usr/local/bin/go-template-engine.zip
RUN unzip go-template-engine.zip
RUN chmod 0755 /usr/local/bin/go-template-engine
RUN mkdir /app
WORKDIR /app
ENTRYPOINT ["/usr/local/bin/go-template-engine"]