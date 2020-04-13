Bitso v3 GO
=============================

[! [Go Report Card](https://goreportcard.com/badge/github.com/webability-go/bitso)](https://goreportcard.com/report/github.com/webability-go/bitso)
[! [GoDoc](https://godoc.org/github.com/webability-go/bitso?status.png)](https://godoc.org/github.com/webability-go/bitso)
[! [GolangCI](https://golangci.com/badges/github.com/webability-go/bitso.svg)](https://golangci.com)

The Bitso package is a complete Bitso V3 API ready to use package in Golang

Manuals are available on godoc.org [![GoDoc](https://godoc.org/github.com/webability-go/bitso?status.png)](https://godoc.org/github.com/webability-go/bitso)

TO DO:
======
Done:
- public API REST

Working:
- Private API REST

Missing:
- REMITTANCE API
- TRANSFER API
- WEBHOOK API
- WEBSOCKET service

Notes on BITSO API bugs:
- On public trades and user trades, the date created_at is NOT an ISO 8601 timestamp (+0000 instead of +00:00 at the end)
- On user trades, "tid" field is string, not long as stated on API Manuals

Version Changes Control
=======================

v0.1.1 - 2020-01-17
- private access services added:
  - BankCodes
  - UserTrades (missing tid parameters)

v0.1.0 - 2020-01-04
- Remastered the full code under a unique API structure
- public access services are finished
- private access is implemented
- private access services working:
  - AccountStatus
  - AccoutBalance
  - Fees

v0.0.2 - 2020-01-03
- public::AvailableBooks() implemented
- public::Tickers() implemented
- public::Ticker(book) implemented

v0.0.1 - 2020-01-01
- First release
- Package skeletton and basic functions
