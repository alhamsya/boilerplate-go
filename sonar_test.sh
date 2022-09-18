#!/bin/sh

sudo rm -rf cover.out
go test -v -short -coverprofile=cover.out `go list ./... | grep -vE '(vendor|mock)'`

exclude=**/*_test.go,**/vendor/**,**/*.js,**/*.jsx,**/*.ts,**/*.html,**/*.css,*mock*.*,**mock**,**/mocks/**,**/files/**,**/dummy/**,**/testgenerator/**,wire.go,wire_gen.go,**/testdata/*,**/cmd/**,**/_*/**,**/models/**,**/model/**
excludeLatest=**/testfile/**,**/docs/**,**/*.java,**/*_test.go,**/vendor/**,**/*.js,**/*.jsx,**/*.ts,**/*.html,**/*.css,**/*.scss,**/*mock*.*,**/mock/**,**/mocks/**,**/files/**,**/dummy/**,**/testgenerator/**,wire.go,wire_gen.go,**/testdata/**,**/cmd/**,**/test/**,**/*.pb.go,**/constant/**,**/constants/**,**/const/**,**/cons/**,**/types/**,**/*types*.go,**/*type*.go,**/*constant*.go,**/*const*.go,**/*cons.go,**/dummy_*.go,**/models/**
sonar-scanner \
  -Dsonar.projectKey=boilerplate-go \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://localhost:9000 \
  -Dsonar.login=218fea8a20094fa3617458e6a45aba85372fcc38 \
  -Dsonar.exclusions=$excludeLatest \
  -Dsonar.go.coverage.reportPaths=cover.out -X