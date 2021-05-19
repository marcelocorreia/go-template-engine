FROM alpine
MAINTAINER marcelo correia

COPY ./bin/go-template-engine /bin/
RUN chmod 0650 /bin/go-template-engine
CMD ["go-template-engine"]