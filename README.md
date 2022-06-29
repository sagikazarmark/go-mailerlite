# Go client for the [MailerLite API](https://developers.mailerlite.com)

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/sagikazarmark/go-mailerlite/CI?style=flat-square)](https://github.com/sagikazarmark/go-mailerlite/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/sagikazarmark/go-mailerlite?style=flat-square)](https://goreportcard.com/report/github.com/sagikazarmark/go-mailerlite)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.18-61CFDD.svg?style=flat-square)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/github.com/sagikazarmark/go-mailerlite)
[![built with nix](https://img.shields.io/badge/builtwith-nix-7d81f7?style=flat-square)](https://builtwithnix.org)

**go-mailerlite is a Go client library for accessing the [MailerLite API v2](https://developers.mailerlite.com).**

**⚠️ WARNING: This is still work in progress. ⚠️**


## Installation

```shell
go get github.com/sagikazarmark/go-mailerlite
```


## API coverage

The following groups of API calls are supported by the client:

- [ ] Campaigns
- [ ] Segments
- [ ] Subscribers
- [ ] Groups
- [ ] Fields
- [ ] Webhooks
- [ ] Stats
- [ ] Settings
- [ ] Batch

Feel free to send PRs to add support for more API calls.


## Development

When all coding and testing is done, please run the test suite:

```shell
make check
```

For the best developer experience, install [Nix](https://builtwithnix.org/) and [direnv](https://direnv.net/).

Alternatively, install Go manually or using a package manager. Install the rest of the dependencies by running `make deps`.
