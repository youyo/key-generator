FROM golang:1.8-jessie
MAINTAINER youyo

ENV APP_DIR /go/src/github.com/youyo/key-generator/

ADD . ${APP_DIR}
WORKDIR ${APP_DIR}
RUN apt-get install make git gcc && \
	make deps

EXPOSE 1323:1323
ENTRYPOINT ["go","run","server.go"]
