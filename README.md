# Code Climate Gofmt Engine

`codeclimate-gofmt` is a Code Climate engine that wraps [Gofmt](https://golang.org/cmd/gofmt/). You can run it on your command line using the Code Climate CLI, or on our hosted analysis platform.

Gofmt automatically formats Go code.

### Installation

1. If you haven't already, [install the Code Climate CLI](https://github.com/codeclimate/codeclimate).
2. Add the following to your Code Climate config:
  ```yaml
  plugins:
    gofmt:
      enabled: true
  ```
3. Run `codeclimate engines:install`
4. You're ready to analyze! Browse into your project's folder and run `codeclimate analyze`.

### Building

```console
make image
```

This will build a `codeclimate/codeclimate-gofmt` image locally.

### Updating

`gofmt` is a part of the Go distribution and shares version with it. Once in a
while a new version of Go gets packaged. In order to get the latest version
and force a new docker image build, please update the base image in the
`Dockerfile`. Please avoid any unstable tags such as `latest` and keep it
explicit.

### Need help?

For help with Gofmt, [check out their documentation](https://golang.org/cmd/gofmt/).

If you're running into a Code Climate issue, first look over this project's [GitHub Issues](https://github.com/codeclimate/codeclimate-rubocop/issues), as your question may have already been covered. If not, [go ahead and open a support ticket with us](https://codeclimate.com/help).
