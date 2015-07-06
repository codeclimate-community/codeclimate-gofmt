# Code Climate Gofmt Engine

`codeclimate-gofmt` is a Code Climate engine that wraps [Gofmt](https://golang.org/cmd/gofmt/). You can run it on your command line using the Code Climate CLI, or on our hosted analysis platform.

Gofmt automatically formats Go code.

### Installation

1. If you haven't already, [install the Code Climate CLI](https://github.com/codeclimate/codeclimate).
2. Run `codeclimate engines:enable gofmt`. This command both installs the engine and enables it in your `.codeclimate.yml` file.
3. You're ready to analyze! Browse into your project's folder and run `codeclimate analyze`.

### Building

In order to build the docker image, you first need to compile a binary for the container. To do that, first [install goxc]() and then run

```console
goxc -bc="linux" -tasks-=go-install,go-vet,go-test,package,archive -d=. -arch=amd64 && mv snapshot/linux_amd64/codeclimate-gofmt bin/codeclimate-gofmt && \
rm -rf snapshot && \
docker build -t codeclimate/codeclimate-gofmt .
```

### Need help?

For help with Gofmt, [check out their documentation](https://golang.org/cmd/gofmt/).

If you're running into a Code Climate issue, first look over this project's [GitHub Issues](https://github.com/codeclimate/codeclimate-rubocop/issues), as your question may have already been covered. If not, [go ahead and open a support ticket with us](https://codeclimate.com/help).
