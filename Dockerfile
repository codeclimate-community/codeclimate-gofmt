FROM golang:alpine

LABEL maintainer="Code Climate <hello@codeclimate.com>"

RUN adduser -u 9000 -D app

WORKDIR /usr/src/app

COPY engine.json /engine.json
COPY codeclimate-gofmt.go ./
RUN apk add --no-cache --virtual .dev-deps git && \
  go get -t -d -v . && \
  go build -o codeclimate-gofmt . && \
  rm -r $GOPATH/src/* && \
  apk del .dev-deps

USER app

VOLUME /code

CMD ["/usr/src/app/codeclimate-gofmt"]
