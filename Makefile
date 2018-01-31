.PHONY: image

IMAGE_NAME ?= codeclimate/codeclimate-gofmt

image:
	docker build --tag "$(IMAGE_NAME)" .
