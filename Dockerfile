FROM yunspace/alpine-golang

ADD build/codeclimate-gofmt /usr/src/app/

RUN adduser -u 9000 -D app
USER app

CMD ["/usr/src/app/codeclimate-gofmt"]
