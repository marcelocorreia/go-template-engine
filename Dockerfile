FROM alpine:3.9

ADD ./bin/go-template-engine /usr/local/bin/go-template-engine
RUN chmod 0755 /usr/local/bin/go-template-engine

ENTRYPOINT ["/usr/local/bin/go-template-engine"]