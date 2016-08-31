FROM golang:alpine

MAINTAINER Code Climate <hello@codeclimate.com>

RUN adduser -u 9000 -D app
USER app

WORKDIR /code

COPY build/codeclimate-gofmt /usr/src/app/

VOLUME /code

CMD ["/usr/src/app/codeclimate-gofmt"]
