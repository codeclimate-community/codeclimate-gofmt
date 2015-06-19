FROM codeclimate/alpine-ruby:0.0.2

WORKDIR /usr/src/app
COPY . /usr/src/app

CMD ["/usr/src/app/bin/codeclimate-gofmt"]
