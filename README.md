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
  - [x] List
  - [ ] Create
  - [x] Get
  - [ ] Update
  - [ ] Search
  - [ ] Search (minimized)
  - [ ] List groups
  - [ ] Activity
  - [ ] Activity (by type)
- [ ] Groups
  - [ ] List
  - [ ] Get
  - [ ] Search
  - [ ] Create
  - [ ] Update
  - [ ] Delete
  - [x] Add subscriber
  - [ ] Import subscribers
  - [ ] Get imports
  - [ ] Add subscriber (with group name)
  - [ ] Assign subscriber
  - [ ] List subscribers
  - [ ] List subscribers (by type)
  - [ ] Get subscriber
  - [ ] Delete subscriber
- [x] Fields
- [ ] Webhooks
- [x] Stats
- [ ] Settings
- [ ] Batch

Feel free to send PRs to add support for more API calls.


## Development

TBD


## License

The project is licensed under the [MIT License](LICENSE).
