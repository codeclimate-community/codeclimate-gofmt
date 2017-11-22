.PHONY: update image

IMAGE_NAME ?= codeclimate/codeclimate-gofmt

update:
	docker run \
	  --rm --interactive \
	  -v $(PWD)/engine.json:/engine.json \
	  -v $(PWD)/bin/update:/usr/local/bin/update \
	  golang:alpine \
	  /usr/local/bin/update

image:
	docker build --tag "$(IMAGE_NAME)" .
