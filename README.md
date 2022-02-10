# <img align="right" src="https://avatars.githubusercontent.com/u/56905970?s=60&v=4" alt="boilerplate-go" title="boilerplate-go" /> boilerplate-go

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/alhamsya/boilerplate-go)
[![Sourcegraph](https://sourcegraph.com/github.com/alhamsya/boilerplate-go/-/badge.svg)](https://sourcegraph.com/github.com/alhamsya/boilerplate-go?badge)
[![Documentation](https://godoc.org/github.com/alhamsya/boilerplate-go?status.svg)](https://godoc.org/github.com/alhamsya/boilerplate-go)
[![codecov](https://codecov.io/gh/alhamsya/boilerplate-go/branch/master/graph/badge.svg?token=mDCHTd8WM7)](https://codecov.io/gh/alhamsya/boilerplate-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/alhamsya/boilerplate-go)](https://goreportcard.com/report/github.com/alhamsya/boilerplate-go)
[![License](https://img.shields.io/github/license/alhamsya/boilerplate-go?color=blue)](https://raw.githubusercontent.com/alhamsya/boilerplate-go/master/LICENSE)

## 👀 Contract API
- [Documentation Postman](https://documenter.getpostman.com/view/2516369/UV5Ro1M1)
- [Download File Collection v2.1 (Postman)](https://github.com/alhamsya/boilerplate-go/blob/master/boilerplate-go.postman_collection.json)
- [Import Collection By Link](https://www.getpostman.com/collections/b71ad1f701723738498f)

## ⚙️ Installation

Make sure you have Go installed. Version `1.14` or higher is required.
Run command for configuration from Makefile 
```bash
make start
```

## ⚙️ Essential Tools gRPC
- https://github.com/protocolbuffers/protobuf/releases/tag/v3.6.1
- https://github.com/protocolbuffers/protobuf-go/releases/tag/v1.27.1

## ⚡️ Quickstart

### Rest
```bash
make run-rest
```

### gRPC
```bash
make run-grpc
```

## 🎯 Features
- Clean Architecture
- Implementation circuit breaker using call wrapper, so that easy to use
- Auto migration from schema directory
- CORS middleware for Fiber that that can be used to enable Cross-Origin Resource Sharing with various options.
- [Limiter middleware](https://github.com/gofiber/fiber/tree/master/middleware/limiter#limiter-middleware)
- DDD (Domain Driven Design) Concept
- Rest API using [Fiber Framework](https://github.com/gofiber/fiber)
- Graceful Handling for gRPC server
- Best Practice Connection Pooling to Database
- Auto migration using Dockerize
- Example code has been implement call to [OMDB API](http://www.omdbapi.com/)
- Database connect loop
- Implementation go-redis
- Handling can be customised by providing an alternate gRCP recovery function
- Integration of logrus logging library into gRPC handlers
- Support development in the Apple Silicon (M1 chipset)