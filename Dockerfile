FROM codeclimate/alpine-ruby:0.0.2

WORKDIR /usr/src/app
COPY . /usr/src/app

RUN adduser -u 9000 -D app
USER app

CMD ["/usr/src/app/bin/codeclimate-gofmt"]
